// queue.go
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQ struct to hold the connection details
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ instance
func NewRabbitMQ(amqpURI string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn, ch}, nil
}

// DeclareExchangeTopicQueue declares exchange, topic, and queue
func (r *RabbitMQ) DeclareExchangeTopicQueue(exchangeName, topicName, queueName string) error {
	err := r.channel.ExchangeDeclare(
		exchangeName, // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		return err
	}

	_, err = r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = r.channel.QueueBind(
		queueName,    // queue name
		topicName,    // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// CloseConnection closes the RabbitMQ connection
func (r *RabbitMQ) CloseConnection() {
	r.channel.Close()
	r.conn.Close()
}
