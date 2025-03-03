/**
 * @nc app=nowcoder id=288815f39cd54b908eed6eae65d69f84 topic=317 question=2563841 lang=Go
 * 2025-02-28 09:16:58
 * https://www.nowcoder.com/practice/288815f39cd54b908eed6eae65d69f84?tpId=317&tqId=2563841
 * [GP40] 绝对值
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return int整型
 */
func absfunc(x int) int {
	// write code here
	if x >= 0 {
		return x
	} else {
		return 0 - x
	}
}

/** @nc code=end */
