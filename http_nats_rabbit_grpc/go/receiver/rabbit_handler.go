package receiver

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *Server) RabbitHandler() {
	msgs := s.consumers["default"]
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
		log.Printf("Received a message: %s", d.Body)
	}
}
