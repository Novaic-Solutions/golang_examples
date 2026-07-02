package main

import (
	"embed"
	"fmt"
	"os"
	//"reflect"
	"gopkg.in/yaml.v3"
	"encoding/json"
)

//go:embed test.yaml
var yamlFile embed.FS

type ApiRequest struct {
	Api_Key string `yaml:"api_key"`
	Api_Secret string `yaml:"api_secret"`
	Base_URL string `yaml:"base_url"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Uri string `yaml:"uri"`
	Method string `yaml:"method"`
	ResponseType string `yaml:"response_type"`
	Params string `yaml:"params"`
	RequestBody any `yaml:"request_body"`
}


func main() {

	file, err := yamlFile.ReadFile("test.yaml")
	if err != nil{
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	apiReq := ApiRequest{}

	if err := yaml.Unmarshal(file, &apiReq); err != nil {
		fmt.Println("Error serializing yaml to object")
		os.Exit(1)
	}

	fmt.Println(apiReq)

	for _, value := range apiReq.Endpoints {
		fmt.Println("---------------------------------")
		//fmt.Println(value)
		//fmt.Println(reflect.TypeOf(value.RequestBody))
		test := value.RequestBody.(map[string]interface{})
		bytes, err := json.Marshal(test)
		if err != nil {
			fmt.Printf("Error marshaling to JSON: %v", err)
		}
		
		fmt.Println(string(bytes))
		
		for k, v := range test {
			fmt.Println("k: ",k, "v :", v)
		}
	}


	// fmt.Println(reflect.TypeOf(apiReq))
	// fmt.Println(apiReq.Body)
	// fmt.Println(reflect.TypeOf(apiReq.Body))

	// //for k, v := range apiReq.Body {
	// //	fmt.Println("k: ",k, "v :", v)
	// //}

	// //fmt.Println(apiReq.Body["current"])

	// jsonBytes, err := json.Marshal(apiReq.Body)
	// if err != nil{
	// 	fmt.Print("Issue serializing into json bytes")
	// 	os.Exit(1)
	// }

	// test := apiReq.Body.(map[string]interface{})
	// fmt.Println(test["current"])
	// fmt.Println(reflect.TypeOf(test["current"]))
	// fmt.Println(reflect.TypeOf(test["resolve"]))

	// fmt.Println(string(jsonBytes))
	// fmt.Println( reflect.TypeOf(jsonBytes))
}

