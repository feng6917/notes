package Q

import (
	"fmt"
	"lgo/QA/circularReference/Ans1/store"
)

type People struct {
}

func (c *People) Money() {
	fmt.Println("客户拿钱 ！！")
}

func (c *People) Buy() {
	fmt.Println("用户购买商品.")
	repo := store.Store{}
	repo.Goods()
	c.Money()
}
