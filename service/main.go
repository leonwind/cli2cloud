package main

import (
	"log"
	"service/api"
)

const (
	port = ":8080"
)

func main() {
	//grpcService := NewService()
	service := api.NewServer()

	if err := service.Start(port); err != nil {
		log.Fatal("Can't start server", err)
	}
}
