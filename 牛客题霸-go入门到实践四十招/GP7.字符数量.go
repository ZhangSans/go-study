/**
 * @nc app=nowcoder id=66f61cc68c2248d08673f2a8df658b94 topic=317 question=2559381 lang=Go
 * 2025-02-24 08:58:03
 * https://www.nowcoder.com/practice/66f61cc68c2248d08673f2a8df658b94?tpId=317&tqId=2559381
 * [GP7] 字符数量
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return int整型
 */
func count(s string) int {
	// write code here
	var array = []rune(s)
	return len(array)
}

/** @nc code=end */
