/**
 * @nc app=nowcoder id=4b05c1f9f0094f86aa14649485e9e9eb topic=317 question=2559448 lang=Go
 * 2025-02-27 14:51:37
 * https://www.nowcoder.com/practice/4b05c1f9f0094f86aa14649485e9e9eb?tpId=317&tqId=2559448
 * [GP29] 字符串构成
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param des string字符串
 * @param src string字符串
 * @return bool布尔型
 */
func canConstruct(des string, src string) bool {
	// write code here
	srcMap := []byte(src)
	desMap := []byte(des)
	have := make(map[byte]int)
	for _, v := range srcMap {
		have[v]++
	}
	for _, v1 := range desMap {
		if have[v1] == 0 {
			return false
		}
		have[v1]--
	}

	return true

}

/** @nc code=end */
