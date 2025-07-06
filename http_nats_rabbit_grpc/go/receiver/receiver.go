package receiver

import (
	"encoding/json"
	"fmt"
	"http_nats_rabbit_grpc/rabbit"
	"io"
	"log"
	"net/http"
)

type ReceiverServerOpts struct {
	Port string
}

type Server struct {
	opts ReceiverServerOpts
}

func StartServerReceiver(opts ReceiverServerOpts) {
	ch := rabbit.ConnectToRabbit()
	msgs, err := ch.Consume(
		"default_queue", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	server := Server{opts: opts}

	mux := http.NewServeMux()
	mux.HandleFunc("/http", server.HttpHandler)

	fmt.Println("receiver before ListenAndServe")
	err = http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	type smallNumber struct {
		One   int8
		Two   int
		Three float32
		Four  float64
	}

	var v smallNumber
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println("cant unmarshal")
	}

	w.Write(body)
}
