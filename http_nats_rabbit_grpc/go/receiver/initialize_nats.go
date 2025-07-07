package receiver

import (
	"log"

	"github.com/nats-io/nats.go"
)

func (s *Server) InitializeNats() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	s.nc = nc
}
