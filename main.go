package main

import (
	"github.com/Novaic-Solutions/golang_examples/client"
	"github.com/Novaic-Solutions/golang_examples/server"
)



func main() {
	incomingChannel := make(chan client.ResponseData, 10)
	server := &server.Server{
		IncomingChannel: incomingChannel,
	}

	c1 := client.NewClient(incomingChannel)
	c1.Start("1")

	c2 := client.NewClient(incomingChannel)
	c2.Start("2")

	server.Start()

}

