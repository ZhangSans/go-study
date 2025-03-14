package main

import "fmt"

type Payload struct {
	Source    string
	TransType string `json:"trans_type"`
	RequestId string `json:"request_id"`
	detect    bool
}

func main() {
	for {
		fmt.Println("请输入要翻译的文本")
		var dirct string
		fmt.Scan(&dirct)

		var url string = "http://api.interpreter.caiyunai.com/v1/translator"
		var token string = "3975l6lr5pcbvidl6jl2"

		var headers map[string]interface{}
		headers["content-type"] = "application/json"
		headers["x-authorization"] = "token " + token

		fmt.Println(dirct)
	}
}
