package todo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DD DbData

type DbData struct {
	BaseDb *gorm.DB
}

// ConnDb 连接数据库
func ConnDb(driver, host string, port int, name, password, dbName, charset, loc string, singular bool, maxIdConn, maxOpenConn int, tableStruct interface{}) (err error) {
	var db *gorm.DB
	db, err = gorm.Open(driver, fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&loc=%s", name, password, host, port, dbName, charset, loc))
	if err != nil {
		return err
	}
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(singular)
	// 设置设置空闲5最大10连接数
	db.DB().SetMaxIdleConns(maxIdConn)
	db.DB().SetMaxOpenConns(maxOpenConn)

	err = db.DB().Ping()
	if err != nil {
		log.Panicf("data ping err: %s", err.Error())
	}
	if !db.HasTable(tableStruct) {
		err = db.AutoMigrate(tableStruct).Error
		if err != nil {
			return err
		}
	} else {
		fmt.Println("has")
		// 自动检查 Product 结构是否变化，变化则进行迁移
		//data.AutoMigrate(table)
	}
	DD.BaseDb = db
	return err
}

// create
func (c *DbData) Create(tableName string, data interface{}) error {
	err := c.BaseDb.Table(tableName).Create(data).Error
	return err
}

// update
func (c *DbData) Update(tableName string, updateParam map[string]interface{}, data interface{}) error {
	err := c.BaseDb.Table(tableName).Where(updateParam).Save(data).Error
	return err
}

// delete
func (c *DbData) Delete(tableName string, deleteParam map[string]interface{}, data interface{}) error {
	err := c.BaseDb.Table(tableName).Delete(data, deleteParam).Error
	return err
}

// get
func (c *DbData) Get(tableName string, getParam map[string]interface{}) (interface{}, error) {
	var data interface{}
	err := c.BaseDb.Table(tableName).Where(getParam).First(&data).Error
	return data, err
}

//func handle(data *gorm.DB,tx *sql.Tx, sqlStatements []string) {
//	var err error
//	if len(sqlStatements) <= 0{
//		return
//	} else{
//		for _, sqlSt := range sqlStatements{
//			err := data.Exec(sqlSt)
//			if err != nil{
//				tx.Rollback()
//				break
//			}
//		}
//	}
//	return
//}

//func (c *DbData) deal(sqlSt []string) handle {
//	data := c.BaseDb
//	tx,err := c.BaseDb.DB().Begin()
//	if err != nil{
//
//	}
//	return handle(func(*data, tx, sqlSt) {
//		return
//	})
//}
//
////  事务
//func (c *DbData) ExecSqlWithTransaction() (err error) {
//	tx, err := c.BaseDb.DB().Begin()
//	if err != nil {
//		return err
//	}
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		}
//	}()
//	if err = handle(tx, sqlStatements); err != nil {
//		tx.Rollback()
//		return err
//	}
//	return tx.Commit()
//}
