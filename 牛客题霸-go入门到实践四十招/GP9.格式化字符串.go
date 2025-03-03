/**
 * @nc app=nowcoder id=00d2408290a94bd5a70965ee0aee2083 topic=317 question=2558985 lang=Go
 * 2025-02-24 09:25:56
 * https://www.nowcoder.com/practice/00d2408290a94bd5a70965ee0aee2083?tpId=317&tqId=2558985
 * [GP9] 格式化字符串
 */

/** @nc code=start */

package main

import (
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 将一个正整数转换为字符串
 * @param num int整型 给定的正整数
 * @return string字符串
 */
func formatstr(num int) string {
	// write code here
	var str1 string = strconv.Itoa(num)
	return str1
}

/** @nc code=end */
