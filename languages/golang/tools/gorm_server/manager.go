package gorm_server

import (
	"errors"
	"fmt"
	"lgo/tools/gorm_server/config"
	"lgo/tools/gorm_server/model"
	"lgo/tools/gorm_server/todo"
)

type Manager struct {
	PeopleRepo model.IPeopleRepo
}

func NewManager(dc *config.DBConfig) (*Manager, error) {
	m := &Manager{}
	var err error
	pr, err := todo.NewPeopleRepo(dc.Driver, dc.Host, dc.Port, dc.Name, dc.Password, dc.Db, dc.Charset, dc.Loc, dc.Singular, dc.MaxIdConn, dc.MaxOpenConn, dc.TablePeople)
	if err != nil {
		err = errors.New(fmt.Sprintf("init people repo err: %s", err.Error()))
		return m, err
	}
	m.PeopleRepo = pr
	return m, err
}
