/**
 * @nc app=nowcoder id=b17a4e37d40549f3ac5ba9ac643d51a5 topic=317 question=2562781 lang=Go
 * 2025-02-27 15:09:04
 * https://www.nowcoder.com/practice/b17a4e37d40549f3ac5ba9ac643d51a5?tpId=317&tqId=2562781
 * [GP30] 不重复的数
 */

/** @nc code=start */

package main

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组
 * @return int整型一维数组
 */
func getNoRepeat(s []int) []int {
	// write code here
	have := make(map[int]int)
	for _, v := range s {
		have[v]++
	}
	var res = make([]int, 0)
	for k, v := range have {
		if v == 1 {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}

/** @nc code=end */
