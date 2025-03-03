/**
 * @nc app=nowcoder id=9baeccf93a58410f9325d70300ce000f topic=317 question=2563964 lang=Go
 * 2025-02-28 11:47:12
 * https://www.nowcoder.com/practice/9baeccf93a58410f9325d70300ce000f?tpId=317&tqId=2563964
 * [GP46] 体温异常
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param t double浮点型 体温
 * @return string字符串
 */
func temperature(t float64) (s string) {
	// write code here
	// if t > 37.5 {
	// 	return errors.New("体温异常").Error()
	// }
	// return ""
	defer func() {
		if err := recover(); err != nil {
			s = "体温异常"
		}
	}()
	if t > 37.5 {
		panic("体温异常")
	} else {
		return s
	}
}

/** @nc code=end */
