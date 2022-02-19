package main

import (
	"log"
	"service/api"
)

const (
	port = ":50051"
)

func main() {
	service := api.NewServer()

	if err := service.Start(port); err != nil {
		log.Fatal("Can't start server", err)
	}
}
