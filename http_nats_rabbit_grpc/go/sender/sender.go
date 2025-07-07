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
}

type Server struct {
	opts SenderServerOpts
	ch   *amqp.Channel
}

func StartServerSender(opts SenderServerOpts) {
	ch := rabbit.ConnectToRabbit()
	server := Server{ch: ch, opts: opts}
	mux := http.NewServeMux()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/rabbit", server.RabbitHandler)

	err := http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func (s *Server) RabbitHandler(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		obj, err := json.Marshal(input)
		if err != nil {
			fmt.Println("cant marshal json")
		}

		err = s.ch.Publish("default_exchange", "default_queue", false, false, amqp.Publishing{Type: "application/json", Body: obj})
		if err != nil {
			panic("sender cant publish")
		}
	}

	w.Write([]byte("done"))
}

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		obj, err := json.Marshal(input)
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

func (s *Server) GetStructByInput() any {
	switch s.opts.TypeOfObjects {
	case "s-number":
		return types.SmallNumber{}
	case "s-string":
		return types.SmallString{}
	case "s-mixed":
		return types.SmallMixed{}
	case "m-number":
		return types.MediumNumber{}
	case "m-string":
		return types.MediumString{}
	case "m-mixed":
		return types.MediumMixed{}
	case "l-number":
		return types.LargeNumber{}
	case "l-string":
		return types.LargeString{}
	case "l-mixed":
		return types.LargeMixed{}
	default:
		fmt.Println(s.opts.TypeOfObjects)
		panic("invalid type sender")
	}
}
