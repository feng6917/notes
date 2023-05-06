package people

import (
	"circularRe/Ans1/Other"
	"fmt"
)

type People struct {
	Other.StoreGoods
}

func (c *People) Money() {
	fmt.Println("客户拿钱 ！！")
}

func (c *People) Buy() {
	fmt.Println("用户购买商品.")
	c.StoreGoods.Goods()
	c.Money()
}
