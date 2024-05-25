package main

import (
	"Chat_App/db"
	"log"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
}
