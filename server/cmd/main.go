package main

import (
	"Chat_App/db"
	"Chat_App/internal/user"
	"Chat_App/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	user.NewService(userRep)
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
