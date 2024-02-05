// main.go
package main

import (
	"fmt"
	"log"
)

func main() {
	amqpURI := "amqp://guest:guest@localhost:5672/"
	exchangeName := "example_exchange"
	topicName := "example.topic"
	queueName := "example_queue"

	rmq, err := NewRabbitMQ(amqpURI)
	if err != nil {
		log.Fatal(err)
	}

	defer rmq.CloseConnection()

	err = rmq.DeclareExchangeTopicQueue(exchangeName, topicName, queueName)
	if err != nil {
		log.Fatal(err)
	}

	// Publish a message
	err = PublishMessage(rmq, exchangeName, topicName, "Hello, RabbitMQ!")
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe and process messages
	msgs, err := Subscribe(rmq, queueName)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		fmt.Printf("Received a message: %s\n", msg.Body)
	}
}
