/**
 * @nc app=nowcoder id=8255c55b41744862bc1cde7ce5b62c42 topic=317 question=2559362 lang=Go
 * 2025-02-27 09:26:42
 * https://www.nowcoder.com/practice/8255c55b41744862bc1cde7ce5b62c42?tpId=317&tqId=2559362
 * [GP22] 评委打分
 */

/** @nc code=start */

package main



/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组 评委给出分数
 * @return int整型一维数组
 */
func minAndMax(s []int) []int {
	// write code here
	var maxScore int
	var minScore int
	for k, v := range s {
		if k == 0 {
			minScore = v
			maxScore = v
		}
		if v >= maxScore {
			maxScore = v
		}
		if v <= minScore {
			minScore = v
		}
	}
	score := make([]int, 0)
	score = append(score, minScore)
	score = append(score, maxScore)
	return score
}

/** @nc code=end */
