/**
 * @nc app=nowcoder id=c44b01f2f525477fb9032c82e8654066 topic=317 question=2563814 lang=Go
 * 2025-02-28 09:05:39
 * https://www.nowcoder.com/practice/c44b01f2f525477fb9032c82e8654066?tpId=317&tqId=2563814
 * [GP39] 数字的阶乘
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param i int整型 数字
 * @return int整型
 */
func factorial(i int) int {
	// write code here
	if i == 0 {
		return 1
	}
	j := factorial(i-1)
	return i * j

}

/** @nc code=end */
