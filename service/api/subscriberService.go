package api

import (
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



// Fetch newest command line output from the server to the live spectator
func SendData(w http.ResponseWriter, request *http.Request) {
	topic := mux.Vars(request)["id"]
	log.Println("Receive messages for topic", topic)
	message := consume(topic, 0)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(message)
}

// consume to a Kafka topic as part of the Publish-Subscriber pattern
func consume(topic string, offset int64) []byte {
	connection, err := sarama.NewConsumer([]string{"kafka:9092"}, sarama.NewConfig())
	if err != nil {
		log.Fatal("Could not connect to kafka:", err)
	}

	partitions, err := connection.Partitions(topic)
	if err != nil {
		log.Fatal("Could not fetch partitions of topic", topic, ":", err)
	}

	consumer, err := connection.ConsumePartition(topic, partitions[0], offset)
	if err != nil {
		log.Fatal("Could not consume partition:", err)
	}

	messages := consumer.Messages()
	var lastMessage []byte
	for msg := range messages {
		lastMessage = msg.Value
		log.Println("Curr message:", msg.Value)
	}
	return lastMessage
}
