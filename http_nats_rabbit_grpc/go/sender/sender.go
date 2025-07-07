package sender

import (
	"http_nats_rabbit_grpc/rabbit"
	"net/http"

	"github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SenderServerOpts struct {
	Port            string
	AmountOfObjects int
	TypeOfObjects   string
}

type Server struct {
	opts SenderServerOpts
	ch   *amqp.Channel
	nc   *nats.Conn
}

func StartServerSender(opts SenderServerOpts) {

	ch := rabbit.ConnectToRabbit()
	server := Server{ch: ch, opts: opts}
	server.InitializeNats()

	mux := http.NewServeMux()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/rabbit", server.RabbitHandler)
	mux.HandleFunc("/nats", server.NatsHandler)

	err := http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}
