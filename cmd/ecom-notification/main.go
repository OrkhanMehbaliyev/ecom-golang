package main

import (
	"context"
	"log"
	"os"

	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-grpc/pb"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-notification/server"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(os.Getenv("SVC_ADDR"), opts...)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewEcomClient(conn)
	srv := server.NewServer(client, &server.AdminInfo{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	})

	done := make(chan struct{})
	go func() {
		srv.Run(context.Background())
		done <- struct{}{}
	}()
	<-done
}
