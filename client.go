package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := client.GetCredentials(ctx, &pb.DBRequest{
		Role: "readonly",
	})

	log.Println("Got credentials:", resp.Username, resp.Password)
}
