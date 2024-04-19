package main

import (
	"fmt"
)

func main() {
	m := Manager{}
	m.Init()
}

type Manager struct {
}

func (c *Manager) Init() {
}

func (c *Manager) Get(name string) {
	fmt.Println("Get: ", name)
}
