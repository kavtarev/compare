package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbit() *amqp.Channel {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Cant obtain channel")

	err = ch.ExchangeDeclare("default_exchange", "direct", false, false, false, false, nil)
	failOnError(err, "Cant create exchange")

	_, err = ch.QueueDeclare("default_queue", false, false, false, false, nil)
	failOnError(err, "Cant create queue")

	err = ch.QueueBind("default_queue", "default_queue", "default_exchange", false, nil)
	failOnError(err, "Cant bind queue")

	return ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
