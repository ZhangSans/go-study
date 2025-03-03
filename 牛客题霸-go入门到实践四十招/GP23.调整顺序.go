/**
 * @nc app=nowcoder id=4a504bffaff74847930c0145886e98ca topic=317 question=2559364 lang=Go
 * 2025-02-27 09:41:15
 * https://www.nowcoder.com/practice/4a504bffaff74847930c0145886e98ca?tpId=317&tqId=2559364
 * [GP23] 调整顺序
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组
 * @return int整型一维数组
 */
func convert(s []int) []int {
	// write code here
	var res []int
	var len int = len(s)
	for i := 0; i < len; i++ {
		res = append(res, s[len-i-1])
	}
	return res
}

/** @nc code=end */
