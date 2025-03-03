/**
 * @nc app=nowcoder id=f6e57de3f6b24220ae4e79ed46ba7ce1 topic=317 question=2559287 lang=Go
 * 2025-02-25 15:05:35
 * https://www.nowcoder.com/practice/f6e57de3f6b24220ae4e79ed46ba7ce1?tpId=317&tqId=2559287
 * [GP14] 器材采购
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型 采购单价
 * @param y int整型 采购单价
 * @param z int整型 采购单价
 * @return int整型
 */
func compare(x int, y int, z int) int {
	// write code here
	if x > y && x > z {
		if y > z {
			return x - z
		} else {
			return x - y
		}
	} else if y > x && y > z {
		if x > z {
			return y - z
		} else {
			return y - x
		}
	} else {
		if x > y {
			return z - y
		} else {
			return z - x
		}
	}
}

/** @nc code=end */
