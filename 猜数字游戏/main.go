package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成随机数
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)
	var number int
	fmt.Println("请输入一个0到100之间的数字")
	// 增加一个标签，跳出双层循环
outloop:
	for {
		fmt.Scan(&number)
		switch {
		case number > target:
			fmt.Println("你猜大了")

		case number < target:
			fmt.Println("你猜小了")
		default:
			fmt.Println("OK，你猜对了")
			break outloop
		}

	}
}
