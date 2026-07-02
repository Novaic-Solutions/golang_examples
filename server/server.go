package server

import (
	"github.com/Novaic-Solutions/golang_examples/client"
	"fmt"
	//"time"
)


type Server struct {
	IncomingChannel chan client.ResponseData
}

func (s *Server) Start() {
	counter := 0
	for {
		fmt.Println("Loop iteration:", counter)

		//---------------------------------------------------------------
		//  By not using the select statement, server will block
		//  until it receives data on the channel.
		//---------------------------------------------------------------
		response := <-s.IncomingChannel
		fmt.Println("Received response at:", response.Timestamp.String(), "Data:", response.Data, "URI:", response.URI)

		//---------------------------------------------------------------
		//	 Using a select statement to listen for incoming data
		//   on the channel allows the server to handle incoming data 
		//	 without blocking the loop.
		//---------------------------------------------------------------

		// select {
		// case response := <-s.IncomingChannel:
		// 	// Process the incoming response data
		// 	// For example, you can log it or perform any other action
		// 	// Here, we just print it to the console
		// 	fmt.Println("Received response at:", response.Timestamp.String(), "Data:", response.Data, "URI:", response.URI)
		// default:
		// 	// No incoming data, you can perform other tasks or sleep for a while
		// 	fmt.Println("No incoming data, sleeping for a while...")
		// 	// Remove the sleep in order to avoid blocking the loop. This will force it to block
		// 	time.Sleep(1 * time.Second)
		// }

		counter++
		if counter >= 25 {
			fmt.Println("Server has processed 25 responses, exiting...")
			return
		}
	}
}