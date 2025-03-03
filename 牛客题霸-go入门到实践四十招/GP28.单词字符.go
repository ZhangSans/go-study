/**
 * @nc app=nowcoder id=a64958131b094faba2adf52a4b131848 topic=317 question=2559377 lang=Go
 * 2025-02-27 10:40:50
 * https://www.nowcoder.com/practice/a64958131b094faba2adf52a4b131848?tpId=317&tqId=2559377
 * [GP28] 单词字符
 */

/** @nc code=start */

package main

import (
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return char字符型
 */
func character(s string) byte {
	// write code here
	count := make(map[byte]int)
	max := 0
	var result byte
	b := []byte(s)

	for _, v := range b {
		count[v]++
		if count[v] > max {
			max = count[v]
			result = v
		}
	}

	return result
}

/** @nc code=end */
