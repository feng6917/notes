package close

import (
	"fmt"
	"time"
)


func CloseA() {
	var ch = make(chan int)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value == 1 {
					fmt.Println("Exit!")
					break
				}
				fmt.Println(value)
			}
		}
	}()

	ch <- 0
	ch <- 1
	close(ch)

	time.Sleep(1 * time.Second)
}

func CloseA1() {
	var ch = make(chan int)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value%9 == 0 {
					fmt.Println("Exit!")
					break
				}
				fmt.Println(value)
			}
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			t := <-ch
			fmt.Println("---- other: ", t)
		}
	}()

	time.Sleep(8 * time.Millisecond)
}

func CloseA2() {
	var ch = make(chan int)
	var chS = make(chan bool)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value%9 == 0 {
					fmt.Println("Exit!")
					chS <- true
					close(chS)
					break
				}
				fmt.Println(value)
			}
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		value := <-chS
		if value {
			close(ch)
		}
	}()

	time.Sleep(1 * time.Second)
}
