/**
 * @nc app=nowcoder id=e385b75257c64296bffa415d3c05fbf0 topic=317 question=2563530 lang=Go
 * 2025-02-28 09:42:37
 * https://www.nowcoder.com/practice/e385b75257c64296bffa415d3c05fbf0?tpId=317&tqId=2563530
 * [GP43] 学生信息III
 */

/** @nc code=start */

package main

import (
	"fmt"
)

type Address struct {
	city     string
	province string
}

type Student struct {
	name    string
	sex     bool
	age     int
	score   int
	address Address
}

func main() {
	student := Student{
		name:  "小明",
		sex:   true,
		age:   23,
		score: 88,
		address: Address{
			city:     "长沙市",
			province: "湖南省",
		},
	}
	fmt.Println(student.name)
	fmt.Println(student.sex)
	fmt.Println(student.age)
	fmt.Println(student.score)
	fmt.Println(student.address.province)
	fmt.Println(student.address.city)
}

/** @nc code=end */
