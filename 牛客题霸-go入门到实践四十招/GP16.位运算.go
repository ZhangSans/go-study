/**
 * @nc app=nowcoder id=c4a2ac91b0bf41529c5be9053f7ba0b9 topic=317 question=2559304 lang=Go
 * 2025-02-25 15:49:56
 * https://www.nowcoder.com/practice/c4a2ac91b0bf41529c5be9053f7ba0b9?tpId=317&tqId=2559304
 * [GP16] 位运算
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型
 * @param b int整型
 * @return int整型一维数组
 */
func bitOperate(a int, b int) []int {
	// write code here
	var slice []int
	slice = append(slice, a&b)
	slice = append(slice, a|b)
	slice = append(slice, a^b)
	return slice
}

/** @nc code=end */
