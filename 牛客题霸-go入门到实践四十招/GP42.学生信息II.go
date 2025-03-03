/**
 * @nc app=nowcoder id=8db171fe5e1542f4a8beaf7268067548 topic=317 question=2563265 lang=Go
 * 2025-02-28 09:29:53
 * https://www.nowcoder.com/practice/8db171fe5e1542f4a8beaf7268067548?tpId=317&tqId=2563265
 * [GP42] 学生信息II
 */

/** @nc code=start */

package main

import (
	"fmt"
)

type Student struct {
	name  string
	sex   bool
	age   int
	score int
}

func main() {
	student := Student{
        name: "小明",
        sex: true,
        age: 23,
        score: 88,
    }
    fmt.Println(student.name)
    fmt.Println(student.sex)
    fmt.Println(student.age)
    fmt.Println(student.score)
}

/** @nc code=end */
