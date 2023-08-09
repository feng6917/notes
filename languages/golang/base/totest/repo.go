package main

import (
	"github.com/jinzhu/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (c *Repo) Init(db *gorm.DB) {
	c.db = db
}

func (c *Repo) GetRepoName() string {
	// 数据操作
	return "repo"
}
