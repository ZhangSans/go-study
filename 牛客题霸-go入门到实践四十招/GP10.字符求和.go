/**
 * @nc app=nowcoder id=79426fd6cf9c49b4828c55d988838372 topic=317 question=2559382 lang=Go
 * 2025-02-24 14:37:25
 * https://www.nowcoder.com/practice/79426fd6cf9c49b4828c55d988838372?tpId=317&tqId=2559382
 * [GP10] 字符求和
 */

/** @nc code=start */

package main

import (
	"fmt"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a string字符串
 * @param b string字符串
 * @return string字符串
 */
func sum(a string, b string) string {
	// write code here
	a1, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Println("转换失败")
	}
	a2, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		fmt.Println("转换失败")
	}
	return strconv.Itoa(int(a1 + a2))
}

/** @nc code=end */
