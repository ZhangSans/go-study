/**
 * @nc app=nowcoder id=022b09f2616a4a69b5930de5bf5151bd topic=317 question=2559379 lang=Go
 * 2025-02-21 16:14:25
 * https://www.nowcoder.com/practice/022b09f2616a4a69b5930de5bf5151bd?tpId=317&tqId=2559379
 * [GP6] 拼接字符串
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串一维数组
 * @return string字符串
 */
func join(s []string) string {
	// write code here
	var str string
	for _, v := range s {

		str = str + v

	}

	return str

}

/** @nc code=end */
