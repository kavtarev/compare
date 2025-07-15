package sender

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) NatsHandler(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		obj, err := json.Marshal(input)
		if err != nil {
			fmt.Println("cant marshal json")
		}
		s.nc.Publish("init", obj)
	}

	w.Write([]byte("done"))
}

func (s *Server) NatsHandlerAutoCannon(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()

	obj, err := json.Marshal(input)
	if err != nil {
		fmt.Println("cant marshal json")
	}

	_, err = s.nc.Request("sync", obj, time.Duration(3*time.Second))
	if err != nil {
		panic(err)
	}

	w.Write([]byte("done"))
}
