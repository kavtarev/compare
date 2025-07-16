package main

import (
	"fmt"
	"http_nats_rabbit_grpc/receiver"
	"http_nats_rabbit_grpc/sender"
)

func main() {
	ch := make(chan int)
	fmt.Println("before servers")

	input := inputCheck()
	senderPort := ":3000"
	receiverPort := ":3001"

	go sender.StartServerSender(sender.SenderServerOpts{
		Port:            senderPort,
		AmountOfObjects: input.numOfRuns,
		TypeOfObjects:   input.jsonType,
	})
	go receiver.StartServerReceiver(receiver.ReceiverServerOpts{
		Port:          receiverPort,
		TypeOfObjects: input.jsonType,
	})

	fmt.Printf("both servers are up on ports sender%v, receiver%v\n", senderPort, receiverPort)

	<-ch
}
