package main

import (
	"context"
	"log"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDBServiceServer
}

func (s *server) GetCredentials(ctx context.Context, req *pb.DBRequest) (*pb.DBResponse, error) {
	log.Println("Received request for role:", req.Role)

	return &pb.DBResponse{
		Username: "user_" + req.Role,
		Password: "pass123",
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	grpcServer := grpc.NewServer()
	pb.RegisterDBServiceServer(grpcServer, &server{})

	log.Println("Spoke plugin running on :50051")
	grpcServer.Serve(lis)
}
