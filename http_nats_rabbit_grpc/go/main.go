package main

import (
	"fmt"
	"http_nats_rabbit_grpc/receiver"
	"http_nats_rabbit_grpc/sender"
)

func main() {
	ch := make(chan int)
	fmt.Println("before servers")

	// input := inputCheck()
	// fmt.Printf("%+v", input)

	go sender.StartServerSender(sender.SenderServerOpts{Port: ":3000"})
	go receiver.StartServerReceiver(receiver.ReceiverServerOpts{Port: ":3001"})

	fmt.Println("after servers")

	<-ch
}
