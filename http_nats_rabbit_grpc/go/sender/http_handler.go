package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		obj, err := json.Marshal(input)
		if err != nil {
			fmt.Println("cant marshal json")
		}
		res, err := http.Post("http://localhost:3001/http", "application/json", bytes.NewBuffer(obj))
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
	}

	w.Write([]byte("done"))
}
