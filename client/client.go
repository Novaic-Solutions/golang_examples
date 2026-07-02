package client

import (
	"time"
	"math/rand/v2"
)

// Will be used for creating a goroutine that sends requests to the api and records responses.

type ResponseData struct {
	Timestamp time.Time
	Data      string
	URI	  string
}

type Client struct {
	ResponseChannel chan ResponseData
}

func NewClient(incomingChannel chan ResponseData) *Client {
	return &Client{
		ResponseChannel: incomingChannel,
	}
}

func (c *Client) Monitor(x string) {
	for {
		result := 1 + rand.IntN(3-1)
		// Simulate sending a request to the API and receiving a response
		response := ResponseData{
			Timestamp: time.Now(),
			Data:      "Sample response data",
			URI:       "Test Endpoint "+x,
		}

		// Send the response data to the channel
		c.ResponseChannel <- response

		// Sleep for a while before sending the next request
		time.Sleep(time.Duration(result) * time.Second)
	}
}

func (c *Client) Start(x string) {
	go c.Monitor(x)
}