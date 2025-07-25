package sender

import (
	"net/http"
	_ "net/http/pprof"
	"time"

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
	opts            SenderServerOpts
	ch              *amqp.Channel
	nc              *nats.Conn
	grpcClient      pb.SenderServiceClient
	totalTime       time.Duration
	ReceivedObjects int
	startTime       time.Time
	endTime         time.Time
	consumers       map[string]<-chan amqp.Delivery
}

func StartServerSender(opts SenderServerOpts) {

	server := Server{opts: opts}
	server.InitializeRabbit()

	// Подключение к gRPC серверу получателя
	conn, err := grpc.NewClient("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go server.RabbitHandlerReceiver()

	server.grpcClient = pb.NewSenderServiceClient(conn)
	server.InitializeNats()

	mux := http.NewServeMux()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/rabbit", server.RabbitHandler)
	mux.HandleFunc("/nats", server.NatsHandler)
	mux.HandleFunc("/grpc", server.GrpcHandler)

	mux.HandleFunc("/grpc-auto-cannon", server.GrpcHandlerAutoCannon)
	mux.HandleFunc("/http-auto-cannon", server.HttpHandlerAutoCannon)
	mux.HandleFunc("/nats-auto-cannon", server.NatsHandlerAutoCannon)

	mux.HandleFunc("/get-time", server.ShowTotalTimeHandler)
	mux.HandleFunc("/get-full-circle-time", server.ShowFullCircleTimeHandler)

	// Добавляем обработчики для pprof
	mux.Handle("/debug/pprof/", http.HandlerFunc(http.DefaultServeMux.ServeHTTP))

	http.ListenAndServe(opts.Port, mux)

}
