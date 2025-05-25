package main

import (
	"log"

	"github.com/OrkhanMehbaliyev/ecom-golang/db"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-api/handler"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-api/server"
	"github.com/OrkhanMehbaliyev/ecom-golang/ecom-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Printf("successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRouters(hdl)
	err = handler.Start(":8080")
	if err != nil {
		log.Fatal("error starting server: %w", err)
	}
}
