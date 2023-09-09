package main

import (
	"log"

	"github.com/evansong007/go-microservices/internal/database"
	"github.com/evansong007/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to inttialize database client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
