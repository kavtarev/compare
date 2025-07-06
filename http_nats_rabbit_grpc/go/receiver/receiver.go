package receiver

import (
	"fmt"
	"http_nats_rabbit_grpc/rabbit"
	"log"
	"net/http"
)

type ReceiverServerOpts struct {
	Port string
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

	mux := http.NewServeMux()
	mux.HandleFunc("/http", HttpHandler)

	fmt.Println("receiver before ListenAndServe")
	err = http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
