package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

// Fetch newest command line output from the server to the live spectator
func SendData(w http.ResponseWriter, request *http.Request) {
	topic := mux.Vars(request)["id"]
	fmt.Println("Receive messages for topic", topic)
	message := consume(topic, 0)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}

// consume to a Kafka topic as part of the Publish-Subscriber pattern
func consume(topic string, offset int64) string {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     topic,
		Partition: 0,
		MinBytes: 1,
		MaxBytes: 100,
	})

	message, err := r.ReadMessage(context.Background())
	if err != nil {
		log.Fatal("Failed to read message:", err)
	}

	fmt.Println("Received message")

	err = r.Close()
	if err != nil {
		log.Fatal("Failed to close reader:", err)
	}

	return string(message.Value)
}
