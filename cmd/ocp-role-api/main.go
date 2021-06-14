package main

import (
	"fmt"

	"google.golang.org/grpc"

	"log"
	"net"

	"github.com/ozoncp/ocp-role-api/internal/api"
	pb "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"
)

const grpcPort = ":82"

func runGrpc() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterOcpRoleApiServer(s, api.NewOcpRoleApi())

	if err := s.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func main() {
	if err := runGrpc(); err != nil {
		log.Fatalf("can't run grpc server: %v", err)
	}
}
