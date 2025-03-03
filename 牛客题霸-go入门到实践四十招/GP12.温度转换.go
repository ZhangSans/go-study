/**
 * @nc app=nowcoder id=3119d996ae674a9fa4649532f76b5d4e topic=317 question=2559253 lang=Go
 * 2025-02-25 14:37:46
 * https://www.nowcoder.com/practice/3119d996ae674a9fa4649532f76b5d4e?tpId=317&tqId=2559253
 * [GP12] 温度转换
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param f double浮点型 华氏温度
 * @return double浮点型
 */
func temperature(f float64) float64 {
	// write code here
	var h float64
	h = (f - 32) * 5 / 9
	return h
}

/** @nc code=end */
