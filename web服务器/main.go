package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建一个http服务器
	http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/get", getHandler)
	// http.HandleFunc("/post", postHandler)
	log.Fatal(http.ListenAndServe(":5012", nil))

}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// func indexHandler(w http.ResponseWriter, req *http.Request) {
// 	var map1 map[string]interface{}
// 	map1["status"] = 200
// 	map1["msg"] = "success"
// 	map1["data"] = nil
// 	data, err := json.Marshal(map1)
// 	if err != nil {
// 		fmt.Println("创建http服务器失败", err)
// 	}
// 	fmt.Println(data)
// }
