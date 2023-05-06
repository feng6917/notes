package other

import (
	"circularRe/Ans2/people"
	"circularRe/Ans2/store"
	"fmt"
)

type Other struct {
	PRepo *people.People
	SRepo *store.Store
}

func (c *Other) Sale() {
	fmt.Println("拿钱给货")
	c.PRepo.Money()
	c.SRepo.Goods()
}

func (c *Other) Buy() {
	fmt.Println("给货拿钱")
	c.SRepo.Goods()
	c.PRepo.Money()
}
