package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		start := time.Now()
		obj, err := json.Marshal(input)
		if err != nil {
			fmt.Println("cant marshal json")
		}
		res, err := http.Post("http://localhost:3001/http", "application/json", bytes.NewBuffer(obj))
		if err != nil {
			panic(err)
		}
		s.totalTime += time.Since(start)
		defer res.Body.Close()
	}

	w.Write([]byte("done"))
}
