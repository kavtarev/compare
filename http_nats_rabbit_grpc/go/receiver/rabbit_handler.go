package receiver

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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

		s.totalTime += time.Since(start)
		log.Printf("Received a message: %s", d.Body)
	}
}
