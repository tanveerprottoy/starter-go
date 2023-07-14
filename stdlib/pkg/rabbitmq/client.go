package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

// "amqp://guest:guest@localhost:5672/"
func (r *RabbitMQ) Connect(url string) {
	var err error
	r.Conn, err = amqp.Dial(url)
	r.handleFailure(err, "Failed to connect to RabbitMQ")
}

func (r *RabbitMQ) Close() {
	r.Conn.Close()
}

func (r *RabbitMQ) handleFailure(err error, msg string) {
	defer r.Close()
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
