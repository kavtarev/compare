package sender

import (
	"http_nats_rabbit_grpc/rabbit"
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
}

func StartServerSender(opts SenderServerOpts) {

	ch := rabbit.ConnectToRabbit()
	server := Server{ch: ch, opts: opts}

	// Подключение к gRPC серверу получателя
	conn, err := grpc.NewClient("localhost:3002", grpc.WithInsecure())
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

	mux.HandleFunc("/get-time", server.ShowTotalTimeHandler)
	mux.HandleFunc("/get-full-circle-time", server.ShowFullCircleTimeHandler)
	mux.HandleFunc("/reset-time", server.ResetTimerHandler)

	// Добавляем обработчики для pprof
	mux.Handle("/debug/pprof/", http.HandlerFunc(http.DefaultServeMux.ServeHTTP))

	http.ListenAndServe(opts.Port, mux)

}
