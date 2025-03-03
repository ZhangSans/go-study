/**
 * @nc app=nowcoder id=1259760c72514e80967e203aa8ce2f81 topic=317 question=2563542 lang=Go
 * 2025-02-28 09:49:52
 * https://www.nowcoder.com/practice/1259760c72514e80967e203aa8ce2f81?tpId=317&tqId=2563542
 * [GP44] 动物和老虎
 */

/** @nc code=start */

package main

import (
	"fmt"
)

type Animaler interface {
	sleep()
	eat()
}
type Tiger struct {
}

func (t Tiger) sleep() {
	fmt.Println("sleep")
}
func (t Tiger) eat() {
	fmt.Println("eat")
}

func main() {
	var t Animaler = Tiger{}
	t.sleep()
	t.eat()
}

/** @nc code=end */
