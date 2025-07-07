package sender

import (
	"encoding/json"
	"fmt"
	"net/http"
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
