/**
 * @nc app=nowcoder id=601a623f5cdb405485fc9ca618e3534d topic=317 question=2564109 lang=Go
 * 2025-02-28 09:00:53
 * https://www.nowcoder.com/practice/601a623f5cdb405485fc9ca618e3534d?tpId=317&tqId=2564109
 * [GP38] 团队闯关
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param score int整型一维数组 团队成员分数
 * @param target int整型 达标分数
 * @return bool布尔型
 */
func canPass(score []int, target int) bool {
	// write code here
	for _, v := range score {
		if v > target {
			return true
		}
	}
	return false
}

/** @nc code=end */
