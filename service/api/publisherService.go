package api

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
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

	log.Println("Publish message:", m.Test, ", for topic:", topic)
	publish(topic, m.Test)
}

// publish the received data to Kafka part of the Publish-subscriber pattern
func publish(topic string, message string) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
	}

	producer, _ := sarama.NewAsyncProducer([]string{"kafka:9092"}, sarama.NewConfig())

	select {
		case producer.Input() <- msg:
			log.Println("Produced message")
		case err := <-producer.Errors():
			log.Println("Failed to produce message:", err)
	}
}
