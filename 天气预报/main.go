package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "net/http"
	"os"
)

func main() {
	// 获取输入的城市名称
	var city string
	fmt.Printf("请输入城市：")
	fmt.Scan(&city)
	// fmt.Println(city)
	// 读取json 匹配城市，获取对应城市code码
	file, err := os.Open("city.json")
	if err != nil {
		fmt.Println("读取城市code码错误", err)
		return
	}
	defer file.Close()
	jsonStr, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读取城市code码错误", err)
		return
	}

	// cityStr := fmt.Sprintf("%s", json)
	// fmt.Println(cityMap)
	var cityMap []map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &cityMap)
	if err != nil {
		fmt.Println("读取城市code码解析json错误", err)
	}
	var nameValue string
	var codeValue string
	for _, v := range cityMap {
		if cityInfo,ok := v.(map[string]interface{});ok{
			
		}


		if cityCode, ok := v["city_code"]; ok {
			codeValue = cityCode
		}
		if cityName, ok := v["city_name"]; ok {
			nameValue = cityName
		}
		if cityName != nil && cityCode != nil {

		}

		// for k, v1 := range v {
		// 	fmt.Println(k,v1)
		// 	break
		// }
		// fmt.Println(v["city_code"])
		// break
	}

	// json = string(json)
	// fmt.Println(string(json))
	// 调用查询天气的api
	// var url string = "http://apis.juhe.cn/simpleWeather/query"
	// var key string = "bef2ae20fc12a6dd10ebc30e0c4ac659"
	// url = url + "?tity=" + city + "&key=" + key
	// res, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("调用天气查询api出错 :", err)
	// 	return
	// }
	// str, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	fmt.Println("解析天气查询结果错误 :", err)
	// 	return
	// }
	// fmt.Printf("%v的天气是%s", city, str)

	// 解析json
	// str, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println("解析天气查询结果错误 :", err)
	// 	return
	// }
	// //输出对应城市的天气
	// fmt.Printf("%v的天气是%v", city, str)
}
