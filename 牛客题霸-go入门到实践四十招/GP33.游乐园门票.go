/**
 * @nc app=nowcoder id=c48714d088954151b0c69ff370e4aee5 topic=317 question=2563998 lang=Go
 * 2025-02-27 15:42:37
 * https://www.nowcoder.com/practice/c48714d088954151b0c69ff370e4aee5?tpId=317&tqId=2563998
 * [GP33] 游乐园门票
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param hight double浮点型 身高
 * @return bool布尔型
 */
func ispay(hight float64) bool {
	// write code here
	if hight >= 160 {
		return true
	}
	return false
}

/** @nc code=end */
