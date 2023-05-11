package other

import (
	"fmt"
	"lgo/QA/circularReference/Ans2/people"
	"lgo/QA/circularReference/Ans2/store"
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
