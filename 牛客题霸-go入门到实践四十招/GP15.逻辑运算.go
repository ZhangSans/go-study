/**
 * @nc app=nowcoder id=00354759fc8049fabf116c5d74a91805 topic=317 question=2559297 lang=Go
 * 2025-02-25 15:33:20
 * https://www.nowcoder.com/practice/00354759fc8049fabf116c5d74a91805?tpId=317&tqId=2559297
 * [GP15] 逻辑运算
 */

/** @nc code=start */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a bool布尔型
 * @param b bool布尔型
 * @return bool布尔型一维数组
 */
func logicalOperation(a bool, b bool) []bool {
	// write code here
	var a1 bool = a && b
	var a2 bool = a || b
	var a3 bool = !a
	var a4 bool = !b
	slice := [4]bool{a1, a2, a3, a4}
	return slice[:]
}

/** @nc code=end */
