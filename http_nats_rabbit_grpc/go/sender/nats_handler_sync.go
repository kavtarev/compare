package sender

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) NatsHandlerSync(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()

	obj, err := json.Marshal(input)
	if err != nil {
		fmt.Println("cant marshal json")
	}
	msg, err := s.nc.Request("sync", obj, time.Duration(3*time.Second))
	if err != nil {
		panic(err)
	}

	w.Write([]byte(msg.Data))
}
