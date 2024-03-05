package main

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	Repo *Repo
}

func (c *Service) Init() *Service {
	gdb, err := gorm.Open("sqlite3", "pwd.db")
	if err != nil {
		panic(err)
	}
	r := Repo{db: gdb}
	s := Service{Repo: &r}
	return &s
}

func (c *Service) GetDaoName() string {
	return c.Repo.GetRepoName()
}
