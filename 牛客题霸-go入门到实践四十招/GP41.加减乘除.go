/**
 * @nc app=nowcoder id=5a21787ce920496a9a165729f5367bd6 topic=317 question=2563921 lang=Go
 * 2025-02-28 09:22:47
 * https://www.nowcoder.com/practice/5a21787ce920496a9a165729f5367bd6?tpId=317&tqId=2563921
 * [GP41] 加减乘除
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
func operate(a int, b int) []int {
	// write code here
	var result []int
	result = append(result, a+b)
	result = append(result, a-b)
	result = append(result, a*b)
	result = append(result, int(a/b))
	result = append(result, a%b)
	return result
}

/** @nc code=end */
