package receiver

import (
	"context"
	"fmt"
	pb "http_nats_rabbit_grpc/grpc"
)

func (s *Server) SendData(ctx context.Context, req *pb.SmallNumber) (*pb.DataResponse, error) {
	fmt.Println("in grpc receiver")
	return &pb.DataResponse{Message: fmt.Sprintf("Received: %+v", req)}, nil
}
