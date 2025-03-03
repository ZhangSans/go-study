/**
 * @nc app=nowcoder id=33dffe16f4554089aa5841b1c49a2bd0 topic=317 question=2563946 lang=Go
 * 2025-02-28 11:42:59
 * https://www.nowcoder.com/practice/33dffe16f4554089aa5841b1c49a2bd0?tpId=317&tqId=2563946
 * [GP45] 网络延迟
 */

/** @nc code=start */

package main

import (
	"errors"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param ping int整型 网络延迟值
 * @return string字符串
 */
func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		return errors.New("网络延迟").Error()
	}
	return ""
}

/** @nc code=end */
