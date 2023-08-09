package main

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	d *Dao
}

func (c *Service) Init(
	:= gorm.Open("sqlite3", "pwd.db")
	if err != nil {
		panic(err)
	}
	d := Dao{}
	d.
	c.d = 
)

func (c *Service) GetDaoName() string {
	return c.d.GetDaoName()
}
