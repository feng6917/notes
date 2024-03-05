package todo

import (
	"errors"
	"fmt"
	"lgo/tools/gorm_server/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// https://learnku.com/docs/gorm/v1/belongs_to/3789

type PeopleRepo struct {
	Db     *gorm.DB
	DbName string
}

func NewPeopleRepo(driver, host string, port int, name, password, dbName, charset, loc string, singular bool, maxIdConn, maxOpenConn int, tablePeople string) (*PeopleRepo, error) {
	var err error
	if DD.BaseDb == nil {
		err = ConnDb(driver, host, port, name, password, dbName, charset, loc, singular, maxIdConn, maxOpenConn, model.People{})
	}
	return &PeopleRepo{Db: DD.BaseDb, DbName: model.People{}.TableName()}, err
}

// create data
func (repo *PeopleRepo) CreateData(p *model.People) (err error) {
	//err = repo.Db.Model(repo.DbName).Create(&p).Error
	err = DD.Create(repo.DbName, &p)
	return err
}

// 获取第一条记录 按主键排序 data.First(&struct)
// 获取一条记录 不指定排序 data.Take(&struct)
// 获取最后一条记录 按照主键排序 data.Last(&struct)
// 获取所有的记录 data.Find(&structs)
// 通过主键进行查询（仅适用于主键是数字类型） data.First(&struct, 10)
// =/<>/in/LIKE/>/</BETWEEN AND/  获取一条或多条
// struct/map/slice 处理 通过 struct 进行查询的时候，GORM 将会查询这些字段的非零值， 意味着你的字段包含 0， ''， false 或者其他 零值 , 将不会出现在查询语句中
// Not 和 Where 查询类似
// 子查询 data.Where("amount > ?", DB.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)
// 指定要从数据库检索的字段  data.Select("")/data.Select([]string{})
// 多个条件排序 data.Order("age desc").Order("name)
// limit offset count group have join

// update data
func (repo *PeopleRepo) UpdateData(p *model.People) error {
	if p.Id <= 0 {
		return errors.New(fmt.Sprintf("id not nil;"))
	}
	err := DD.Update(repo.DbName, map[string]interface{}{"id": p.Id}, &p)
	return err
}

// data.Save(&struct) 更新所有字段
// update 更新单个字段
// updates map[string]interface{} 只更新修改的字段
// updates struct 只更新修改的和非空的字段
// 只想更新或者忽略某些字段，可以使用 Select，Omit 方法。
// 带表达式更新
//

// delete data
func (repo *PeopleRepo) DeletePeople(deleteParam map[string]interface{}) error {
	if deleteParam == nil {
		return nil
	} else {
		fmt.Printf("data:%+v\r\n", deleteParam)
		err := DD.Delete(repo.DbName, deleteParam, &model.People{})
		return err
	}
}

// get people

func (repo *PeopleRepo) GetPeople(getParam map[string]interface{}) (interface{}, error) {
	var data interface{}
	var err error
	if getParam == nil {
		return data, nil
	} else {
		data, err = DD.Get(repo.DbName, getParam)
		return data, err
	}
}
