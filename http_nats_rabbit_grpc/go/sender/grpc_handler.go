package sender

import (
	"context"
	"fmt"
	pb "http_nats_rabbit_grpc/grpc"
	"net/http"
	"time"
)

func (s *Server) GrpcHandler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		start := time.Now()
		req := &pb.LargeMixed{}
		_, err := s.grpcClient.SendData(context.Background(), req)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to send data", http.StatusInternalServerError)
			return
		}
		s.totalTime += time.Since(start)
	}
	fmt.Fprintf(w, "Response from receiver: %s")
}
