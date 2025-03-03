/**
 * @nc app=nowcoder id=0ca84c08e2874e31983eb61875158099 topic=317 question=2559350 lang=Go
 * 2025-02-26 15:59:09
 * https://www.nowcoder.com/practice/0ca84c08e2874e31983eb61875158099?tpId=317&tqId=2559350
 * [GP19] 创建切片
 */

/** @nc code=start */

package main

import ()

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param length int整型 切片初始化长度
 * @param capacity int整型 切片初始化容量
 * @return int整型一维数组
 */
func makeslice(length int, capacity int) []int {
	// write code here
	var slice1 = make([]int, length, capacity)
	for i := 0; i < length; i++ {
		slice1[i] = i
	}
	return slice1
}

/** @nc code=end */
