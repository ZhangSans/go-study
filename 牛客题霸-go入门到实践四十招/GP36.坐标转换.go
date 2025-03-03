/**
 * @nc app=nowcoder id=948b9ca358c146dabe3e7d30b1d5e34c topic=317 question=2559340 lang=Go
 * 2025-02-28 08:49:05
 * https://www.nowcoder.com/practice/948b9ca358c146dabe3e7d30b1d5e34c?tpId=317&tqId=2559340
 * [GP36] 坐标转换
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型二维数组
 * @return int整型一维数组
 */
func transform(array [][]int) []int {
	// write code here
	var map1 []int
	for _, v := range array {
		for _, vv := range v {
			map1 = append(map1, vv)

		}
	}
	return map1
}

/** @nc code=end */
