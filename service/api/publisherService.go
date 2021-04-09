package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

type msg struct {
	Test string
}

// Receive data from the cmd line
func ReceiveData(w http.ResponseWriter, request *http.Request) {
	topic := mux.Vars(request)["id"]

	var m msg
	err := json.NewDecoder(request.Body).Decode(&m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("Publish message:", m.Test, ", for topic:", topic)
	publish(topic, m.Test)
}

// publish the received data to Kafka part of the Publish-subscriber pattern
func publish(topic string, message string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "new-topic",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key-b"),
			Value: []byte(message),
		})

	if err != nil {
		log.Fatal("Failed to publish message:", err)
	}

	err = w.Close()
	if err != nil {
		log.Fatal("Failed to close writer:", err)
	}
}
