package receiver

import (
	"context"
	"fmt"
	pb "http_nats_rabbit_grpc/grpc"
)

func (s *Server) SendData(ctx context.Context, req *pb.LargeMixed) (*pb.DataResponse, error) {
	return &pb.DataResponse{Message: fmt.Sprintf("Received: %+v", req)}, nil
}
