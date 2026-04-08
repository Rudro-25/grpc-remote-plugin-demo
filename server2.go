package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-demo/proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDBServiceServer
	db *sql.DB
}

func (s *server) GetCredentials(ctx context.Context, req *pb.DBRequest) (*pb.DBResponse, error) {
	username := fmt.Sprintf("user_%d", time.Now().Unix())
	password := "pass123"

	query := fmt.Sprintf(
		"CREATE USER %s WITH PASSWORD '%s';",
		username, password,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return nil, err
	}

	log.Println("Created DB user:", username)

	return &pb.DBResponse{
		Username: username,
		Password: password,
	}, nil
}

func main() {
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDBServiceServer(grpcServer, &server{db: db})

	log.Println("Spoke plugin (DB-enabled) running on :50051")
	grpcServer.Serve(lis)
}
