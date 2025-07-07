package receiver

import (
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
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
}

func StartServerReceiver(opts ReceiverServerOpts) {

	server := Server{opts: opts}
	mux := http.NewServeMux()

	server.InitializeRabbit()
	go server.RabbitHandler()

	server.InitializeNats()
	server.NatsHandler()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/get-time", server.ShowTotalTimeHandler)
	mux.HandleFunc("/reset-time", server.ResetTimerHandler)

	err := http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}
