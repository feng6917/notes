package people

import (
	"fmt"
)

type People struct {
}

func (c *People) Money() {
	fmt.Println("客户拿钱 ！！")
}
