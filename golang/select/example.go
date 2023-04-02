package c

import (
	"golang/select/sbreak"
)

func mBreak() {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	//sbreak.BreakA(ch)
	//sbreak.BreakA1(ch)

	sbreak.BreakA2(ch)
}
