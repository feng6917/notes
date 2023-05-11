package Q

import (
	"fmt"
	"lgo/QA/circularReference/Ans1/people"
)

type Store struct {
}

func (c *Store) Goods() {
	fmt.Println("拿货！！")
}

func (c *Store) Sale() {
	fmt.Println("商店出售商品.")
	repo := people.People{}
	repo.Money()
	c.Goods()
}
