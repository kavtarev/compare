package sender

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"http_nats_rabbit_grpc/rabbit"
	"net/http"
)

type SenderServerOpts struct {
	Port string
}

func StartServerSender(opts SenderServerOpts) {
	ch := rabbit.ConnectToRabbit()

	err := ch.Publish("default_exchange", "default_queue", false, false, amqp.Publishing{Type: "application/json", Body: []byte("sender")})
	if err != nil {
		panic("sender cant publish")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/http", HttpHandler)

	fmt.Println("sender before ListenAndServe")
	err = http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("some")
	w.Write([]byte("hello"))
}
