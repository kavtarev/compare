package receiver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "http_nats_rabbit_grpc/grpc"
)

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &pb.LargeMixed{})
	if err != nil {
		fmt.Println("cant unmarshal")
	}

	w.Write(body)
}
