package receiver

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func (s *Server) NatsHandler() {
	s.nc.Subscribe("init", func(m *nats.Msg) {
		into := s.GetStructByInput()
		start := time.Now()

		err := json.Unmarshal(m.Data, &into)
		if err != nil {
			fmt.Println("cant unmarshal")
		}
		s.totalTime += time.Since(start)

		m.Respond([]byte("answer is 42"))
	})
}
