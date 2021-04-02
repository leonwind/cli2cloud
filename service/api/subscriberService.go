package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"time"
)

// Fetch newest command line output from the server to the live spectator
func SendData(w http.ResponseWriter, request *http.Request) {
	topic := mux.Vars(request)["id"]
	message := consume(topic, 10)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}

// consume to a Kafka topic as part of the Publish-Subscriber pattern
func consume(topic string, partition int) string {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetReadDeadline(time.Now().Add(10*time.Second))
	if err != nil {
		log.Fatal("Failed to set read deadline: ", err)
	}

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	message := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(message)
		if err != nil {
			break
		}
		fmt.Println(string(message))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

	return message
}
