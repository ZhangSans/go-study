/**
 * @nc app=nowcoder id=572280082f414dde9207e11790e823d5 topic=317 question=2562806 lang=Go
 * 2025-02-24 09:05:28
 * https://www.nowcoder.com/practice/572280082f414dde9207e11790e823d5?tpId=317&tqId=2562806
 * [GP8] 回文数
 */

/** @nc code=start */

package main

import (
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	// write code here
	var arr string = strconv.Itoa(x)
	length := len(arr)
	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-i-1] {
			return false
		}
	}
	return true
}

/** @nc code=end */
