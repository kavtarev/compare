package receiver

import (
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *Server) RabbitHandler() {
	msgs := s.consumers["default_queue"]
	for d := range msgs {
		start := time.Now()
		into := s.GetStructByInput()
		err := json.Unmarshal(d.Body, &into)
		if err != nil {
			fmt.Println("cant unmarshal")
		}

		obj, err := json.Marshal(into)
		if err != nil {
			fmt.Println("receiver cant marshal json")
		}

		s.totalTime += time.Since(start)
		_, err = s.ch.PublishWithDeferredConfirm("default_exchange", "default_queue_response", false, false, amqp.Publishing{Type: "application/json", Body: obj})
		if err != nil {
			panic("receiver cant publish")
		}
		d.Ack(false)
	}
}
