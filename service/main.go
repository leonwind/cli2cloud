package main

import (
	"log"
	"service/api"
)

const (
	port = ":50051"
)

func main() {
	service, err := api.NewServer()
	if err != nil {
		log.Fatal("Cant create server", err)
	}

	if err := service.Start(port); err != nil {
		log.Fatal("Can't start server", err)
	}
}
