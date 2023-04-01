package store

import (
	"fmt"
)

type Store struct {
}

func (c *Store) Goods() {
	fmt.Println("拿货！！")
}
