package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 获取输入的城市名称
	var city string
	fmt.Printf("请输入城市：")
	fmt.Scan(&city)
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
	var cityCode string
	for _, v := range cityMap {
		if v["city_name"] == city {
			if code, ok := v["city_code"].(string); ok {
				cityCode = code
				break
			}
		}
	}
	if cityCode == "" {
		fmt.Printf("%v 未到找对应的城市编码，请换个城市", city)
	}
	// fmt.Println(cityCode)

	// 调用查询天气的api
	var url string = "http://t.weather.sojson.com/api/weather/city/"
	url = url + cityCode
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("调用天气查询api出错 :", err)
		return
	}
	str, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("解析天气查询结果错误 :", err)
		return
	}
	// fmt.Printf("%v的天气是%s", city, str)
	var response map[string]interface{}
	err = json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Println("解析城市天气数据json出现错误", err)
		return
	}
	if response["status"] != float64(200) {
		fmt.Println("查询天气数据出现错误", response["message"])
		return
	}
	// //输出对应城市的天气
	fmt.Printf("城市: %v", city)
	fmt.Printf("时间: %v", response["time"])
	//类型转换
	data, ok := response["data"].(map[string]interface{})
	if ok {
		fmt.Printf("湿度: %v\n", data["shidu"])
		fmt.Printf("pm25: %v\n", data["pm25"])
		fmt.Printf("pm10: %v\n", data["pm10"])
		fmt.Printf("空气质量: %v\n", data["quality"])
		fmt.Printf("温度: %v\n\r\n", data["wendu"])
		forecast, ok := data["forecast"].([]interface{})
		if ok {
			for _, v := range forecast {
				v1, ok := v.(map[string]interface{})
				if ok {
					fmt.Printf("日期: %v\n", v1["ymd"])
					fmt.Printf("星期: %v\n", v1["week"])
					fmt.Printf("风向: %v\n", v1["fx"])
					fmt.Printf("风力: %v\n", v1["fl"])
					fmt.Printf("天气: %v\n\r\n", v1["type"])
				}

			}
		}
	}
}
