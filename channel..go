package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := "your_exchange"
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare exchange")

	topics := []string{"topic.one", "topic.two", "topic.three"}

	for _, topic := range topics {
		q, err := ch.QueueDeclare(
			"",    // name (let RabbitMQ generate a unique name)
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = ch.QueueBind(
			q.Name,        // queue name
			topic,         // routing key
			exchangeName,  // exchange
			false,         // no-wait
			nil,           // arguments
		)
		failOnError(err, "Failed to bind a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		go func(topic string, msgs <-chan amqp.Delivery) {
			for d := range msgs {
				fmt.Printf("Received a message from topic %s: %s\n", topic, d.Body)
			}
		}(topic, msgs)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
