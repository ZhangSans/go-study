package main

import "fmt"

func main() {
	for {
		fmt.Println("请输入要翻译的文本")
		var dirct string
		fmt.Scan(&dirct)
		fmt.Println(dirct)
	}
}
