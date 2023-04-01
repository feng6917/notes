package store

import (
	"circularRe/Ans1/Other"
	"fmt"
)

type Store struct {
	Other.PeopleMoney
}

func (c *Store) Goods() {
	fmt.Println("拿货！！")
}

func (c *Store) Sale() {
	fmt.Println("商店出售商品.")
	c.PeopleMoney.Money()
	c.Goods()
}
