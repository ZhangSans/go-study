/**
 * @nc app=nowcoder id=7526b138abee44be927a535787172b8b topic=317 question=2564105 lang=Go
 * 2025-02-28 08:53:45
 * https://www.nowcoder.com/practice/7526b138abee44be927a535787172b8b?tpId=317&tqId=2564105
 * [GP37] 质量检查
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param material int整型一维数组 成品质量
 * @param standard int整型 质检标准
 * @return int整型一维数组
 */
func check(material []int, standard int) []int {
	// write code here
	var pass []int
	for _, v := range material {
		if v >= standard {
			pass = append(pass, v)
		}
	}
	return pass
}

/** @nc code=end */
