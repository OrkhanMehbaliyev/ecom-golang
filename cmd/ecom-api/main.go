package main

import (
	"log"
	"os"

	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-api/handler"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-grpc/pb"
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
	hdl := handler.NewHandler(client, os.Getenv("SECRET_KEY"))
	handler.RegisterRouters(hdl)

	err = handler.Start(":8080")
	if err != nil {
		log.Fatal("error starting server: %w", err)
	}
}
