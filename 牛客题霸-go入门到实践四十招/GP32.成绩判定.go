/**
 * @nc app=nowcoder id=b2b3fa396af04925adb5f43653da267c topic=317 question=2559336 lang=Go
 * 2025-02-27 15:31:37
 * https://www.nowcoder.com/practice/b2b3fa396af04925adb5f43653da267c?tpId=317&tqId=2559336
 * [GP32] 成绩判定
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param score int整型 分数
 * @return string字符串
 */
func judgeScore(score int) string {
	// write code here
	var res string
	switch {
	case score < 60:
		res = "不及格"
	case score >= 60 && score < 80:
		res = "中等"
	case score >= 80 && score < 90:
		res = "良好"
	default:
		res = "优秀"
	}
	return res
}

/** @nc code=end */
