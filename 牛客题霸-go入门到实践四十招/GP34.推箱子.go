/**
 * @nc app=nowcoder id=2b11fcf1926744889a71e72d1fbf7f89 topic=317 question=2564099 lang=Go
 * 2025-02-27 15:45:23
 * https://www.nowcoder.com/practice/2b11fcf1926744889a71e72d1fbf7f89?tpId=317&tqId=2564099
 * [GP34] 推箱子
 */

/** @nc code=start */

package main



/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param forwards string字符串 推箱子方向
 * @return bool布尔型
 */
func pushBox(forwards string) bool {
	// write code here

	var x int
	var y int

	for _, v := range forwards {
		switch v {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		default:
			y--
		}
	}
	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}

}

/** @nc code=end */
