package sender

import (
	"encoding/json"
	"fmt"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *Server) RabbitHandler(w http.ResponseWriter, r *http.Request) {
	input := s.GetStructByInput()
	for i := 0; i < s.opts.AmountOfObjects; i++ {
		obj, err := json.Marshal(input)
		if err != nil {
			fmt.Println("cant marshal json")
		}

		_, err = s.ch.PublishWithDeferredConfirm("default_exchange", "default_queue", false, false, amqp.Publishing{Type: "application/json", Body: obj})

		if err != nil {
			panic("sender cant publish")
		}
	}

	w.Write([]byte("done"))
}
