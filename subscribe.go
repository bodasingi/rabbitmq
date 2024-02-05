// subscribe.go
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Subscribe subscribes to a queue and processes messages
func Subscribe(rmq *RabbitMQ, queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := rmq.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
