package sender

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (s *Server) RabbitHandlerReceiver() {
	msgs := s.consumers["default_queue_response"]
	for d := range msgs {
		s.ReceivedObjects++
		if s.ReceivedObjects == 1 {
			s.startTime = time.Now()
		}
		start := time.Now()
		into := s.GetStructByInput()
		err := json.Unmarshal(d.Body, &into)
		if err != nil {
			fmt.Println("cant unmarshal")
		}
		d.Ack(false)
		if s.ReceivedObjects == s.opts.AmountOfObjects {
			s.endTime = time.Now()
		}

		s.totalTime += time.Since(start)

		log.Printf("Received a message: %s", d.Body)
	}
}
