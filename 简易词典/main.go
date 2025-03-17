package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Source    string
	TransType string `json:"trans_type"`
	RequestId string `json:"request_id"`
	Detect    bool
}

func main() {
	for {
		fmt.Println("请输入要翻译的文本")
		var dirct string
		fmt.Scan(&dirct)

		var url string = "http://api.interpreter.caiyunai.com/v1/translator"
		var token string = "3975l6lr5pcbvidl6jl2"

		requestMap := map[string]interface{}
		requestMap["source"] = dirct
		requestMap["trans_type"] = "auto2zh"
		requestMap["request_id"] = "demo"
		requestMap["detect"] = true

		requestBody, err := json.Marshal(requestMap)

		var headers map[string]interface{}
		headers["content-type"] = "application/json"
		headers["x-authorization"] = "token " + token

		payload := Payload{
			Source:    dirct,
			TransType: "auto2zh",
			RequestId: "demo",
			Detect:    true,
		}
		http.Post(url)

		fmt.Println(dirct)
	}
}
