package store

import (
	"fmt"
	"lgo/questionWithAnswer/circularReference/answerOne/other"
)

type Store struct {
	other.PeopleMoney
}

func (c *Store) Goods() {
	fmt.Println("拿货！！")
}

func (c *Store) Sale() {
	fmt.Println("商店出售商品.")
	c.PeopleMoney.Money()
	c.Goods()
}
