package main

import (
	"fmt"
	"sync"
)

var num int

var mu sync.Mutex

func Add(wg *sync.WaitGroup) {
	mu.Lock()
	num += 1
	mu.Unlock()
	wg.Done()
}

func main() {

	var a, b int
	if a&b != 1{
		fmt.Println("hello")
	}
	c := 1
	cn := c|a
	fmt.Println(cn)

	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go Add(wg)
	}

	wg.Wait()
	println(num) // Output: 100

}
