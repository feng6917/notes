package model

import (
	"github.com/jinzhu/gorm"
	"time"
)


type People struct {
	Id        uint32 `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Name      string `json:"name" gorm:"column:name;type:varchar(256);not null;default ''"`
	Password  string `json:"password" gorm:"column:password;type:varchar(256);not null;default ''"`
	CreatedAt uint64 `json:"created_at" gorm:"column:created_at;not null;default 0"`
	UpdatedAt uint64 `json:"updated_at" gorm:"column:updated_at;not null;default 0"`
}

func (People) TableName() string {
	return "gorm_people"
}

// 在钩子中设置字段值
func (c *People) BeforeCreate(scope *gorm.Scope) {
	c.Id = uint32(time.Now().Unix())
	c.CreatedAt = uint64(time.Now().Unix())
	c.UpdatedAt = c.CreatedAt
	scope.SetColumn("id", uint32(time.Now().Unix()))
	scope.SetColumn("created_at", c.CreatedAt)
	scope.SetColumn("updated_at", c.UpdatedAt)
}


// 在钩子中设置字段值
func (c *People) BeforeUpdate(scope *gorm.Scope){
	c.UpdatedAt = uint64(time.Now().Unix())
	scope.SetColumn("updated_at", c.UpdatedAt)
}

type IPeopleRepo interface {
	CreateData(p *People) error
	UpdateData(p *People) error
	DeletePeople(deleteParam map[string]interface{}) error
	GetPeople(getParam map[string]interface{}) (interface{}, error)
}
