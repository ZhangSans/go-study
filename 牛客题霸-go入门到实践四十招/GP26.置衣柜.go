/**
 * @nc app=nowcoder id=3f9beccc1cc443afa4c32dde9e23bb2e topic=317 question=2564112 lang=Go
 * 2025-02-27 10:09:44
 * https://www.nowcoder.com/practice/3f9beccc1cc443afa4c32dde9e23bb2e?tpId=317&tqId=2564112
 * [GP26] 置衣柜
 */

/** @nc code=start */

package main

import  "fmt"

func main() {
  var slice1 =make([]string, 0)
  //解法1
//   slice1 = append(slice1, "帽子")
//   slice1 = append(slice1, "围巾")
//   slice1 = append(slice1, "衣服")
//   slice1 = append(slice1, "裤子")
//   slice1 = append(slice1, "袜子")
  //解法2
  slice1 = append(slice1, "帽子","围巾","衣服","裤子","袜子")
  fmt.Println(slice1)
}

/** @nc code=end */
