package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"http_nats_rabbit_grpc/rabbit"
	"http_nats_rabbit_grpc/types"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SenderServerOpts struct {
	Port            string
	AmountOfObjects int
	TypeOfObjects   string
	SizeOfObjects   string
}

type Server struct {
	opts SenderServerOpts
	ch   *amqp.Channel
}

func StartServerSender(opts SenderServerOpts) {
	ch := rabbit.ConnectToRabbit()

	server := Server{
		ch:   ch,
		opts: opts,
	}

	err := ch.Publish("default_exchange", "default_queue", false, false, amqp.Publishing{Type: "application/json", Body: []byte("sender")})
	if err != nil {
		panic("sender cant publish")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/http", server.HttpHandler)

	fmt.Println("sender before ListenAndServe")
	err = http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		obj, err := json.Marshal(types.SmallNumber{})
		if err != nil {
			fmt.Println("cant marshal json")
		}
		res, err := http.Post("http://localhost:3001/http", "application/json", bytes.NewBuffer(obj))
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

	}

	w.Write([]byte("done"))
}
