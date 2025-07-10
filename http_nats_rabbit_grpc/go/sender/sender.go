package sender

import (
	"http_nats_rabbit_grpc/rabbit"
	"net/http"

	pb "http_nats_rabbit_grpc/grpc"

	"github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type SenderServerOpts struct {
	Port            string
	AmountOfObjects int
	TypeOfObjects   string
}

type Server struct {
	opts       SenderServerOpts
	ch         *amqp.Channel
	nc         *nats.Conn
	grpcClient pb.SenderServiceClient
}

func StartServerSender(opts SenderServerOpts) {

	ch := rabbit.ConnectToRabbit()
	server := Server{ch: ch, opts: opts}

	// Подключение к gRPC серверу получателя
	conn, err := grpc.NewClient("localhost:3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	server.grpcClient = pb.NewSenderServiceClient(conn)
	server.InitializeNats()

	mux := http.NewServeMux()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/rabbit", server.RabbitHandler)
	mux.HandleFunc("/nats", server.NatsHandler)
	mux.HandleFunc("/grpc", server.GrpcHandler)

	err = http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}
