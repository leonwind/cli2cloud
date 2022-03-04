package main

import (
	"log"
	"service/api"
)

const (
	port  = ":50051"
	dbUrl = "postgres://leon.windheuser@localhost:5432/cli2cloud"
)

func main() {
	service, err := api.NewServer(dbUrl)
	if err != nil {
		log.Fatal("Cant create server", err)
	}

	if err := service.Start(port); err != nil {
		log.Fatal("Can't start server", err)
	}
}
