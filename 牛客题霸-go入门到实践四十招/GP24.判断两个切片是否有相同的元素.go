/**
 * @nc app=nowcoder id=9fb825cc67514400ae6ba7d72db3541c topic=317 question=2562785 lang=Go
 * 2025-02-27 09:49:51
 * https://www.nowcoder.com/practice/9fb825cc67514400ae6ba7d72db3541c?tpId=317&tqId=2562785
 * [GP24] 判断两个切片是否有相同的元素
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s1 int整型一维数组
 * @param s2 int整型一维数组
 * @return bool布尔型
 */
func equal(s1 []int, s2 []int) bool {
	// write code here
	var equal bool = true
	if len(s1) == len(s2) {
		for i, v := range s1 {
			if v != s2[i] {
				equal = false
				break
			}
		}
	} else {
		equal = false
	}
	return equal
}

/** @nc code=end */
