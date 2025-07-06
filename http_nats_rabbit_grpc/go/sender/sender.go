package sender

import (
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SenderServerOpts struct {
	Port string
}

func StartServerSender(opts SenderServerOpts) {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Cant obtain channel")

	err = ch.ExchangeDeclare("default", "direct", false, false, false, false, nil)
	failOnError(err, "Cant create exchange")

	_, err = ch.QueueDeclare("default", false, false, false, false, nil)
	failOnError(err, "Cant create queue")

	fmt.Println(conn)

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
