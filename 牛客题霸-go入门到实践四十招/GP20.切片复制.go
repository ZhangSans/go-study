/**
 * @nc app=nowcoder id=007731bde92a4291a2fe10ef2bc57cea topic=317 question=2559352 lang=Go
 * 2025-02-26 16:06:28
 * https://www.nowcoder.com/practice/007731bde92a4291a2fe10ef2bc57cea?tpId=317&tqId=2559352
 * [GP20] 切片复制
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param src int整型一维数组 源切片
 * @param des int整型一维数组 目的切片
 * @return int整型一维数组
 */
func sliceCopy(src []int, des []int) []int {
	// write code here
	des = src
	return des
}

/** @nc code=end */
