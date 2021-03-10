package main

import (
	api "github.com/leonwind/cli2cloud/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
