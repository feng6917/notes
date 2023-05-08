package goroutine

import (
	"fmt"
	"sync"
)

func mWaitGroup() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 2000; i > 0; i-- {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
