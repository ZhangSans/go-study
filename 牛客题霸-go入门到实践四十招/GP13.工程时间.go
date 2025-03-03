/**
 * @nc app=nowcoder id=9b39f10933f146faa193a4cb9654b067 topic=317 question=2559269 lang=Go
 * 2025-02-25 14:48:00
 * https://www.nowcoder.com/practice/9b39f10933f146faa193a4cb9654b067?tpId=317&tqId=2559269
 * [GP13] 工程时间
 */

/** @nc code=start */

package main

import (
	"fmt"
)

func main() {

	var totalDay int
	totalDay = 97
	var week int = totalDay / 7
	var day int = totalDay % 7
	fmt.Println(week)
	fmt.Println(day)
}

/** @nc code=end */
