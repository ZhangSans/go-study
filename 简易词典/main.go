package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for {
		fmt.Println("请输入要翻译的文本")
		var dirct string
		fmt.Scan(&dirct)

		var url string = "http://api.interpreter.caiyunai.com/v1/translator"
		var token string = "3975l6lr5pcbvidl6jl2"

		requestMap := map[string]interface{}{
			"source":     dirct,
			"trans_type": "auto2zh",
			"request_id": "demo",
			"detect":     true,
		}

		requestBody, err := json.Marshal(requestMap)
		if err != nil {
			fmt.Println("requestMap json序列化失败", err)
			return
		}
		//创建一个新的请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("创建request请求对象失败 ", err)
			return
		}
		//设置header头
		req.Header.Set("content-type", "application/json")
		req.Header.Set("x-authorization", "token "+token)
		//发送请求
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("发送请求失败", err)
			return
		}
		defer response.Body.Close()
		str, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("读取请求数据失败", err)
			return
		}
		// fmt.Println(response.Status)
		var responseMap map[string]interface{}
		err = json.Unmarshal([]byte(str), &responseMap)
		if err != nil {
			fmt.Println("解析请求数据失败", err)
			return
		}
		fmt.Println("翻译结果：", responseMap["target"])
		fmt.Sprintf("\r\n")

		// fmt.Println(string(str))
	}
}
