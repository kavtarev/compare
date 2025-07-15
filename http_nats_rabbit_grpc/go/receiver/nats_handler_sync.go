package receiver

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func (s *Server) NatsHandlerSync() {
	s.nc.Subscribe("sync", func(m *nats.Msg) {
		fmt.Println("in receiver sync")
		into := s.GetStructByInput()
		start := time.Now()

		err := json.Unmarshal(m.Data, &into)
		if err != nil {
			fmt.Println("cant unmarshal")
		}
		s.totalTime += time.Since(start)

		m.RespondMsg(m)
		// m.Respond([]byte("answer is 42"))
	})
}
