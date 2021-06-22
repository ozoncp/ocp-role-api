package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"log"
	"net"

	"github.com/ozoncp/ocp-role-api/internal/api"
	"github.com/ozoncp/ocp-role-api/internal/producer"
	"github.com/ozoncp/ocp-role-api/internal/repo"
	"github.com/ozoncp/ocp-role-api/internal/tracing"

	pb "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	grpcPort         = ":8082"
	dbHost           = "localhost"
	dbPort           = 5432
	dbUser           = "postgres"
	dbPassword       = "postgres"
	dbName           = "ocp_role_test"
	serviceName      = "ocp-role-api"
	kafkaHostPort    = "127.0.0.1:9094"
	kafkaTopic       = "test"
	producerBuffSize = 10
)

func createDBConnect() (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"pgx",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName,
		),
	)

	return db, err
}

func runGrpc(db *sqlx.DB) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()

	producer, err := NewKafkaProducer(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create kafka producer: %w", err)
	}
	defer producer.Close()

	pb.RegisterOcpRoleApiServer(s, api.NewOcpRoleApi(repo.New(db), producer))

	if err := s.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()
	err := pb.RegisterOcpRoleApiHandlerFromEndpoint(
		ctx, rmux, "localhost"+grpcPort, []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.WriteHeader(200)
		http.ServeFile(w, r, "./swagger/api.swagger.json")
	})
	mux.Handle("/", rmux)

	return http.ListenAndServe("localhost:8080", mux)
}

func NewKafkaProducer(ctx context.Context) (producer.Producer, error) {
	p, err := producer.NewProducer([]string{kafkaHostPort}, kafkaTopic)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka conn: %w", err)
	}

	buffered, err := producer.NewBuffered(p, producerBuffSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka conn: %w", err)
	}

	go func() {
		for {
			select {
			case <-buffered.Done():
				return
			case e := <-buffered.C():
				err := p.Send(e)
				if err != nil {
					log.Printf("kafka producer send error: %v", err)
				}
			}
		}
	}()

	return buffered, nil
}

func main() {
	tracing.InitTracing(serviceName)

	go func() {
		db, err := createDBConnect()
		if err != nil {
			log.Fatalf("fail to open DB connection: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("fail to ping DB: %v", err)
		}

		defer func() {
			if err := db.Close(); err != nil {
				log.Printf("failed to close DB connection: %v", err)
			}
		}()

		if err := runGrpc(db); err != nil {
			log.Fatalf("can't run grpc server: %v", err)
		}
	}()

	if err := runHTTP(); err != nil {
		log.Fatalf("can't run HTTP server: %v", err)
	}
}
