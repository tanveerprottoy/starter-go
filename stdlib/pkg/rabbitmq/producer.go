package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
}

func (p *Producer) Send(conn *amqp.Connection) {
	ch, err := conn.Channel()
	p.handleFailure(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	p.handleFailure(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	p.handleFailure(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func (p *Producer) handleFailure(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
