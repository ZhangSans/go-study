/**
 * @nc app=nowcoder id=db550f43f4d244b2ad9f630880f756b5 topic=317 question=2559327 lang=Go
 * 2025-02-27 15:25:41
 * https://www.nowcoder.com/practice/db550f43f4d244b2ad9f630880f756b5?tpId=317&tqId=2559327
 * [GP31] 年龄分段
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param age int整型 年龄
 * @return string字符串
 */
func getAge(age int) string {
	// write code here
	if age >= 0 && age < 1 {
		return "婴儿"
	} else if age >= 1 && age <= 4 {
		return "幼儿"
	} else if age >= 5 && age <= 11 {
		return "儿童"
	} else if age >= 12 && age <= 18 {
		return "少年"
	} else if age >= 19 && age <= 35 {
		return "青年"
	} else if age >= 36 && age <= 59 {
		return "中年"
	} else {
		return "老年"
	}
}

/** @nc code=end */
