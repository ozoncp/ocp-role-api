package main

import (
	"fmt"

	"google.golang.org/grpc"

	"log"
	"net"

	"github.com/ozoncp/ocp-role-api/internal/api"
	"github.com/ozoncp/ocp-role-api/internal/repo"
	pb "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	grpcPort   = ":8082"
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "test"
	dbName     = "ocp_role_test"
)

func runGrpc() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	db, err := sqlx.Open(
		"pgx",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName,
		),
	)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close DB connection: %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("fail to open DB connection: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOcpRoleApiServer(s, api.NewOcpRoleApi(repo.New(db)))

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
