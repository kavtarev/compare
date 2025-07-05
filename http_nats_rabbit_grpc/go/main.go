package main

import (
	"http_nats_rabbit_grpc/sender"
)

func main() {
	sender.StartServer(sender.SenderServerOpts{Port: ":3000"})
}
