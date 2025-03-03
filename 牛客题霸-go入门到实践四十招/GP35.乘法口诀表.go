/**
 * @nc app=nowcoder id=d4e76eafc39148c3a27af51b4cf2aa5c topic=317 question=2559331 lang=Go
 * 2025-02-27 16:59:00
 * https://www.nowcoder.com/practice/d4e76eafc39148c3a27af51b4cf2aa5c?tpId=317&tqId=2559331
 * [GP35] 乘法口诀表
 */

/** @nc code=start */

package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j)
		}
		fmt.Println()
	}
}

/** @nc code=end */
