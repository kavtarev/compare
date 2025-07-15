package sender

import (
	"http_nats_rabbit_grpc/rabbit"

	"github.com/rabbitmq/amqp091-go"
)

func (s *Server) InitializeRabbit() {
	ch := rabbit.ConnectToRabbit()
	s.ch = ch

	msgs, err := s.ch.Consume(
		"default_queue_response", // queue
		"",                       // consumer
		false,                    // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		panic(err)
	}
	s.consumers = make(map[string]<-chan amqp091.Delivery)

	s.consumers["default_queue_response"] = msgs
}
