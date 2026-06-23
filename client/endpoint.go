package client

import (
	"crypto/tls"
	"net/http"
	"bytes"
	"time"
)

// type Endpoint interface {
// 	SendRequest() *http.Response
// }

type ApiRequest struct {
	Method string
	Endpoint string
	Body []byte
	Username string
	Password string

}

// Create TLS Config
func CreateTlsConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

// Create Transport for HTTP Client
func CreateHTTPTransport() *http.Transport{
	return &http.Transport{
		TLSClientConfig: CreateTlsConfig(),
	}
}

// Create HTTP Client
func CreateHTTPClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: CreateHTTPTransport(),
	}
}

// Create Request with Basic Auth and send it
func SendRequest(apiRequest *ApiRequest) (*http.Response, error) {
	req, err := http.NewRequest(apiRequest.Method, apiRequest.Endpoint, bytes.NewBuffer(apiRequest.Body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiRequest.Username, apiRequest.Password)

	client := CreateHTTPClient(30)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
