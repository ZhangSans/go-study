/**
 * @nc app=nowcoder id=fc88ae20a01e4ff9be14927c53dd6b8a topic=317 question=2559358 lang=Go
 * 2025-02-26 16:31:48
 * https://www.nowcoder.com/practice/fc88ae20a01e4ff9be14927c53dd6b8a?tpId=317&tqId=2559358
 * [GP21] 出队
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组 身高
 * @param index int整型 出队索引
 * @return int整型一维数组
 */
func deleteElement(s []int, index int) []int {
	// write code here
	//方法1
	// var slice = make([]int, len(s)-1)
	// i := 0
	// for k, v := range s {
	// 	if k == index {
	// 		continue
	// 	}
	// 	slice[i] = v
	// 	i++
	// }
	// return slice
	//方法2
	s1 := s[0:index]
	s2 := s[index+1:]
	s1 = append(s1, s2...)
    return s1;
}

/** @nc code=end */
