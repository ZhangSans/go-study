/**
 * @nc app=nowcoder id=994a90bc0e3c4e8a9650b7e55ebd8590 topic=317 question=2562801 lang=Go
 * 2025-02-27 09:59:23
 * https://www.nowcoder.com/practice/994a90bc0e3c4e8a9650b7e55ebd8590?tpId=317&tqId=2562801
 * [GP25] 合并有序数组
 */

/** @nc code=start */

package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums1 int整型一维数组
 * @param m int整型
 * @param nums2 int整型一维数组
 * @param n int整型
 * @return int整型一维数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	for i, v := range nums2 {
        nums1[n+i] =v
	}
	return nums1
}

/** @nc code=end */
