package sender

import (
	"context"
	"fmt"
	pb "http_nats_rabbit_grpc/grpc"
	"net/http"
	"time"
)

func (s *Server) GrpcHandler(w http.ResponseWriter, r *http.Request) {
	s.startTime = time.Now()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		start := time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		_, err := s.grpcClient.SendData(ctx, &pb.LargeMixed{})
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to send data", http.StatusInternalServerError)
			return
		}
		s.ReceivedObjects++
		if s.ReceivedObjects == s.opts.AmountOfObjects {
			s.endTime = time.Now()
			fmt.Println("finally")
		}
		s.totalTime += time.Since(start)
	}
	fmt.Fprintf(w, "Response from receiver: %s")
}

func (s *Server) GrpcHandlerAutoCannon(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := s.grpcClient.SendData(ctx, &pb.LargeMixed{})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to send data", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))
}
