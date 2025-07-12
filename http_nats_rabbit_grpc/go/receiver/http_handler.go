package receiver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pb "http_nats_rabbit_grpc/grpc"
)

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &pb.LargeMixed{})
	if err != nil {
		fmt.Println("cant unmarshal")
	}

	s.totalTime += time.Since(start)
	w.Header().Add("content-type", "application/json")
	w.Write(body)
}
