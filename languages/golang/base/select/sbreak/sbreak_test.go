package sbreak

import "testing"

func TestBreakA(t *testing.T) {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	BreakA(ch)
}

func TestBreakA1(t *testing.T) {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	BreakA1(ch)
}

func TestBreakA2(t *testing.T) {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	BreakA2(ch)
}
