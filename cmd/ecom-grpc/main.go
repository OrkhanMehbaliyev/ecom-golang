package main

import (
	"log"
	"net"
	"os"

	"github.com/OrkhanMehbaliyev/ecom-golang/db"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-grpc/pb"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-grpc/server"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-grpc/storer"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Printf("successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)

	grpcSrv := grpc.NewServer()
	pb.RegisterEcomServer(grpcSrv, srv)

	listener, err := net.Listen("tcp", os.Getenv("SVC_ADDR"))
	if err != nil {
		log.Fatalf("listener failed: %v", err)
	}

	log.Printf("server listening on %s", os.Getenv("SVC_ADDR"))
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
