package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"time"
)

// Receive data from the cmd line
func ReceiveData(w http.ResponseWriter, request *http.Request) {
	topic := mux.Vars(request)["id"]

	var message string;
	err := json.NewDecoder(request.Body).Decode(&message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	publish(topic, 10, message)
}

// publish the received data to Kafka part of the Publish-subscriber pattern
func publish(topic string, partition int, message string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("Failed to dial leader: ", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	if err != nil {
		log.Fatal("Failed to set write Deadline: ", err)
	}

	_, err = conn.WriteMessages(kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Fatal("Failed to write message: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Failed to close connection: ", err)
	}
}