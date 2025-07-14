package receiver

import (
	"net"
	"net/http"
	"time"

	pb "http_nats_rabbit_grpc/grpc"

	"github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type ReceiverServerOpts struct {
	Port          string
	TypeOfObjects string
}

type Server struct {
	opts      ReceiverServerOpts
	ch        *amqp.Channel
	totalTime time.Duration
	nc        *nats.Conn
	consumers map[string]<-chan amqp.Delivery
	pb.UnimplementedSenderServiceServer
}

func StartServerReceiver(opts ReceiverServerOpts) {

	server := Server{opts: opts}
	mux := http.NewServeMux()

	server.InitializeRabbit()
	go server.RabbitHandler()

	server.InitializeNats()
	server.NatsHandler()
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterSenderServiceServer(grpcServer, &server)
		lis, err := net.Listen("tcp", ":3002")
		if err != nil {
			panic(err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/get-time", server.ShowTotalTimeHandler)

	err := http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}
