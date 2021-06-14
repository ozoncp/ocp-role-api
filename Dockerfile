FROM ubuntu AS builder

RUN apt update -y
RUN apt upgrade -y

RUN apt install -y locales
RUN apt install -y sudo

RUN echo "LC_ALL=en_US.UTF-8" >> /etc/environment && \
    echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen && \
    echo "LANG=en_US.UTF-8" > /etc/locale.conf && \
    locale-gen en_US.UTF-8

RUN useradd -m -G sudo developer
RUN echo 'developer:developer' | chpasswd
USER developer

RUN echo developer | sudo -S apt install -y curl
RUN echo developer | sudo -S DEBIAN_FRONTEND="noninteractive" \
    bash -c 'curl -O https://dl.google.com/go/go1.16.linux-amd64.tar.gz && ls -l && \
    tar xvf go1.16.linux-amd64.tar.gz && \
    chown -R developer:developer ./go && \
    mv ./go /usr/local'
RUN echo developer | sudo -S apt install -y ca-certificates && sudo update-ca-certificates
RUN echo developer | sudo -S apt install -y make git vim protobuf-compiler

ENV GOPATH /home/developer/go
ENV PATH $PATH:/usr/local/go/bin:/home/developer/go/bin

COPY . /home/developer/go/src/github.com/ozoncp/ocp-role-api
RUN echo developer | sudo -S chown -R developer /home/developer/

WORKDIR /home/developer/go/src/github.com/ozoncp/ocp-role-api

RUN make install-deps && make gen && make build

# FROM alpine:latest
FROM ubuntu
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/developer/go/src/github.com/ozoncp/ocp-role-api/bin/ocp-role-api .
RUN chown root:root ocp-role-api
EXPOSE 82
CMD ["./ocp-role-api"]
