package main

import (
	"log"

	"github.com/OrkhanMehbaliyev/ecom-golang/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Printf("successfully connected to database")
}
