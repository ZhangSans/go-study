/**
 * @nc app=nowcoder id=0efdfc6bced2443c81857324b6b2f0ad topic=317 question=2559371 lang=Go
 * 2025-02-27 10:36:11
 * https://www.nowcoder.com/practice/0efdfc6bced2443c81857324b6b2f0ad?tpId=317&tqId=2559371
 * [GP27] 成绩表
 */

/** @nc code=start */

package main

import (
	"fmt"
)

func main() {
	var map1 = make(map[string]int, 6)
	map1["小明"] = 60
	map1["小王"] = 70
	map1["张三"] = 95
	map1["张伟"] = 88
	map1["李四"] = 98
	map1["王五"] = 100
	fmt.Println(map1)
}

/** @nc code=end */
