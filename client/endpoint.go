package client

import (
	"crypto/tls"
	"net/http"
	"fmt"
	"bytes"
	"time"
	"io"
)

// type Endpoint interface {
// 	SendRequest() *http.Response
// }

// type ApiResponseData interface {
// 	Serialize(*http.Response) ([]byte, error)
// }

type ApiRequest struct {
	Method string
	Endpoint string
	Body []byte
	Username string
	Password string
}

type ApiResponse struct {
	Endpoint string
	StatusCode int
	Body []byte
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

	fmt.Println("Sending request to:", apiRequest.Endpoint)
	fmt.Println("Request method:", apiRequest.Method)
	fmt.Println("Body: ", apiRequest.Body)

	client := CreateHTTPClient(30)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Create ApiResponse from http.Response json
func CreateApiResponse(resp *http.Response, endpoint string) (*ApiResponse, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		return nil, err
	}

	apiResponse := &ApiResponse{
		Endpoint: endpoint,
		StatusCode: resp.StatusCode,
		Body: body,
	}

	return apiResponse, nil
}
