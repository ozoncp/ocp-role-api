.PHONY: all .vendor-proto check build empty gen install-deps swagger

empty: ;

swagger:
	docker run -p 80:8080 \
	-e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json \
	-v `pwd`/swagger:/swagger swaggerapi/swagger-ui

build:
	go build -o bin/ocp-role-api ./cmd/ocp-role-api

check: pkg/ocp-role-api/*.go internal/*/*.go cmd/*/*.go
	go build  ./internal/*/ ./cmd/*/

gen: pkg/ocp-role-api/*.go
	go1.16.3 generate internal/mockgen.go


.vendor-proto: \
  vendor.protogen/google \
  vendor.protogen/github.com/envoyproxy \
  vendor.protogen/api/ocp-role-api/ocp-role-api.proto ;

vendor.protogen/google:
	git clone --depth=1 https://github.com/googleapis/googleapis vendor.protogen/googleapis && \
	mkdir -p  vendor.protogen/google/ && \
	mv vendor.protogen/googleapis/google/api vendor.protogen/google && \
	rm -rf vendor.protogen/googleapis

vendor.protogen/github.com/envoyproxy:
	mkdir -p vendor.protogen/github.com/envoyproxy && \
	git clone --depth=1 https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate

vendor.protogen/api/ocp-role-api/ocp-role-api.proto: api/ocp-role-api/ocp-role-api.proto
	mkdir -p vendor.protogen/api/ocp-role-api
	cp api/ocp-role-api/ocp-role-api.proto vendor.protogen/api/ocp-role-api/

pkg/ocp-role-api/*.go: vendor.protogen/api/ocp-role-api/ocp-role-api.proto
	mkdir -p pkg/ocp-role-api
	mkdir -p swagger
	protoc \
		-I vendor.protogen \
		--go_out=pkg/ocp-role-api \
		--go_opt=paths=import \
		--go-grpc_out=pkg/ocp-role-api \
		--go-grpc_opt=paths=import \
		--grpc-gateway_out=pkg/ocp-role-api \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=import \
		--validate_out lang=go:pkg/ocp-role-api \
		--swagger_out=allow_merge=true,merge_file_name=api:swagger \
		api/ocp-role-api/ocp-role-api.proto
	
	mv pkg/ocp-role-api/github.com/ozoncp/ocp-role-api/pkg/ocp-role-api/*.go pkg/ocp-role-api/
	rm -rf pkg/ocp-role-api/github.com

install-deps:
	ls go.mod || go mod init
	go get github.com/envoyproxy/protoc-gen-validate
	go install github.com/envoyproxy/protoc-gen-validate
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/envoyproxy/protoc-gen-validate
