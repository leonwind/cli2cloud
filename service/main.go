package main

import (
	"github.com/gorilla/mux"
	"github.com/leonwind/cli2cloud/api"
	"log"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/ping", api.Ping).Methods("GET")
	apiRouter.HandleFunc("/new", api.CreateNewID).Methods("GET")
	// receive the newest command line updates
	apiRouter.HandleFunc("/data/{id}", api.ReceiveData).Methods("POST")
	// send the latest stored updates
	apiRouter.HandleFunc("/data/{id}", api.SendData).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
