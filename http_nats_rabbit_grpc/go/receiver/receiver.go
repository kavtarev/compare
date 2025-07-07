package receiver

import (
	"encoding/json"
	"fmt"
	"http_nats_rabbit_grpc/rabbit"
	"http_nats_rabbit_grpc/types"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

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
	server := Server{opts: opts, ch: ch}
	mux := http.NewServeMux()

	go func() {
		for d := range msgs {
			start := time.Now()
			into := server.GetStructByInput()
			err = json.Unmarshal(d.Body, &into)
			if err != nil {
				fmt.Println("cant unmarshal")
			}

			server.totalTime += time.Since(start)
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	mux.HandleFunc("/http", server.HttpHandler)
	mux.HandleFunc("/get-time", server.ShowTotalTimeHandler)
	mux.HandleFunc("/reset-time", server.ResetTimerHandler)

	err = http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func (s *Server) HttpHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	into := s.GetStructByInput()
	err = json.Unmarshal(body, &into)
	if err != nil {
		fmt.Println("cant unmarshal")
	}

	s.totalTime += time.Since(start)
	w.Header().Add("content-type", "application/json")
	w.Write(body)
}

func (s *Server) ShowTotalTimeHandler(w http.ResponseWriter, r *http.Request) {
	totalTimeStr := strconv.FormatInt(s.totalTime.Microseconds(), 10)
	w.Write([]byte(totalTimeStr))
}

func (s *Server) ResetTimerHandler(w http.ResponseWriter, r *http.Request) {
	s.totalTime = 0
	w.Write([]byte("reset to 0"))
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
		panic("invalid type")
	}
}
