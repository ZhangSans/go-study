/**
 * @nc app=nowcoder id=632df72cb8934791a3a5f873f6306e47 topic=317 question=2563981 lang=Go
 * 2025-02-21 14:35:44
 * https://www.nowcoder.com/practice/632df72cb8934791a3a5f873f6306e47?tpId=317&tqId=2563981
 * [GP5] 值和指针
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型 变量a
 * @param b int整型 变量b
 * @return bool布尔型一维数组
 */
func equal(a int, b int) []bool {
	// write code here
	var1 := make([]bool, 0)
	if &a == &b {
		var1 = append(var1, true)
	} else {
		var1 = append(var1, false)
	}
	if a == b {
		var1 = append(var1, true)

	} else {
		var1 = append(var1, false)
	}
	return var1

}

/** @nc code=end */
