// publish.go
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// PublishMessage publishes a message to a topic
func PublishMessage(rmq *RabbitMQ, exchangeName, routingKey, message string) error {
	err := rmq.channel.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
