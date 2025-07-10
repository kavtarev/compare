package sender

import (
	"context"
	"fmt"
	pb "http_nats_rabbit_grpc/grpc"
	"net/http"
)

func (s *Server) GrpcHandler(w http.ResponseWriter, r *http.Request) {
	req := &pb.SmallNumber{}
	res, err := s.grpcClient.SendData(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to send data", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Response from receiver: %s", res.Message)
}
