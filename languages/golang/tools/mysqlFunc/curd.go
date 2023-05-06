package mysql_func

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"reflect"
)

const (
	user     = "root"
	password = "mysql"
	dbname   = "mysql_demo"
)

var GlobalDb *gorm.DB

func init() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname))
	if err != nil {
		panic(err)
	}
	GlobalDb = db
	GlobalDb.LogMode(true)
	//defer data.Close()
}

// 创建数据
func Create(obj interface{}, tb string) error {
	return GlobalDb.Table(tb).Create(obj).Error
}

// 修改数据
func Update(whereObj, updateObj interface{}, tb string) error {
	// 1. 处理where 条件值
	var db *gorm.DB
	var err error
	db, err = GetUpdateWhere(whereObj, tb)
	if err != nil {
		return err
	}
	// 2. 处理update 条件值
	var up map[string]interface{}
	up, err = GetUpdateUpdate(updateObj)
	if err != nil {
		return err
	}
	log.Printf("update map: %+v\r\n", up)
	// 3. 更新
	return db.Updates(up).Error
}

// 获取数据
func List(whereObj interface{}, tb string, otherObj interface{}) (v []interface{}, total int64, err error) {
	// 1. 处理where 条件值
	var db *gorm.DB
	fmt.Println(whereObj)
	db, err = GetUpdateWhere(whereObj, tb)
	if err != nil {
		return v, total, err
	}
	//dbCErr := data.Count(&total).Error
	//if dbCErr != nil {
	//	return v, total, dbCErr
	//}
	dbDErr := db.Find(&v).Error
	if dbDErr != nil {
		return v, total, dbDErr
	}
	fmt.Println("total: ", total)
	fmt.Println("data: ", v)
	return v, total, err
}

// 获取修改数据where 条件
func GetUpdateWhere(obj interface{}, tb string) (*gorm.DB, error) {
	// 1. 生成id
	var id string
	var err error
	id, err = GenDataId(obj)
	if err != nil {
		return nil, err
	}
	db := GlobalDb.Table(tb)
	// 2. 获取id结构
	fieldList := GlobalDC[id]
	fmt.Println("fl: ", fieldList)
	// 3. 处理数据
	if len(fieldList) > 0 {
		sv := reflect.ValueOf(obj)
		for flIndex, fl := range fieldList {
			if fl != nil {
				var v interface{}
				if fl.Index < sv.NumField() {
					v = sv.Field(fl.Index).Interface()
				}
				// 处理单个字段
				key, val, err := DealSingleField(fl, v)
				if err != nil {
					return db, err
				}
				fieldList[flIndex].Key = key
				fieldList[flIndex].Value = val
				fieldList[flIndex].Sql = fmt.Sprintf("%s %s %s", key, fl.KeyValue, val)
			}
		}
	}
	fmt.Println(fieldList)

	// 4. 生成where data
	andSql, orSql := FieldList2Sql(fieldList)
	if andSql != "" {
		db = db.Where(andSql)
	}
	if orSql != "" {
		db = db.Or(orSql)
	}
	return db, nil
}

// 获取修改数据update 数据
func GetUpdateUpdate(obj interface{}) (map[string]interface{}, error) {
	// 1. 生成id
	var id string
	var err error
	id, err = GenDataId(obj)
	if err != nil {
		return nil, err
	}
	up := make(map[string]interface{}, 0)
	// 2. 获取id结构
	fieldList := GlobalDC[id]
	// 3. 处理数据
	if len(fieldList) > 0 {
		sv := reflect.ValueOf(obj)
		for _, fl := range fieldList {
			var v interface{}
			if fl.Index < sv.NumField() {
				v = sv.Field(fl.Index).Interface()
			}
			// 处理单个字段
			key, val, err := DealSingleFieldWithUpdate(fl, v)
			if err != nil {
				return up, err
			}
			if !fl.IsIgnore && (fl.IsMust || (key != "" && val != "")) {
				up[key] = val
			}
		}
	}
	return up, err

}

// -> andSql; orSql
func FieldList2Sql(fl []*FieldData) (string, string) {
	var andStr, orStr string
	if len(fl) > 0 {
		for _, f := range fl {
			if !f.IsIgnore && (f.IsMust || (f.Sql != "")) {
				if !f.IsOr {
					if andStr != "" {
						andStr += "and "
					}
					andStr += f.Sql
				} else {
					if orStr != "" {
						orStr += "and "
					}
					orStr += f.Sql
				}
			}
		}
	}
	return andStr, orStr
}

// 处理单个字段 获取key value
func DealSingleField(fl *FieldData, value interface{}) (string, string, error) {
	var key, val string
	// 校验是否忽略
	var err error
	if fl.IsIgnore {
		return key, val, err
	}
	// 校验是否必填
	if !fl.IsMust {
		// 校验是否空值校验
		if fl.IsNull {
			// 根据字段类型处理
			if reflect.ValueOf(value).IsZero() {
				return key, val, errors.New(fmt.Sprintf("Err: %v -> not nil", fl.Key))
			}
		}
	}

	// 处理字段key 重置
	if fl.RenameKey != "" {
		key = fmt.Sprintf("`%s`", fl.RenameKey)
	} else {
		key = fmt.Sprintf("`%s`", fl.Key)
	}
	// 根据对应关系处理 value
	if fl.KeyValue == FieldKeyValueEqual ||
		fl.FieldType == FieldKeyValueNotEqual ||
		fl.FieldType == FieldKeyValueMoreThan ||
		fl.FieldType == FieldKeyValueMoreEqual ||
		fl.FieldType == FieldKeyValueLessThan ||
		fl.FieldType == FieldKeyValueLessEqual {
		if fl.FieldType == FieldTypeString {
			val = fmt.Sprintf("'%v'", value)
		} else if fl.FieldType == FieldTypeInt || fl.FieldType == FieldTypeBool {
			val = fmt.Sprintf("%v", value)
		}
	} else if fl.KeyValue == FieldKeyValueIn || fl.KeyValue == FieldKeyValueNotIn {
		if fl.FieldType == FieldTypeIntArray {
			val, err = Interface2ArrayStr(value, true)
			if err != nil {
				return key, val, err
			}
		} else if fl.FieldType == FieldTypeStringArray {
			val, err = Interface2ArrayStr(value, false)
			if err != nil {
				return key, val, err
			}
		}
	} else if fl.KeyValue == FieldKeyValueLike {
		if fl.FieldType == FieldTypeString || fl.FieldType == FieldTypeInt {
			val = fmt.Sprintf("'%%v%'", value)
		}
	}
	return key, val, err
}

// 处理单个字段 获取key value
func DealSingleFieldWithUpdate(fl *FieldData, value interface{}) (string, interface{}, error) {
	var key string
	// 校验是否忽略
	var err error
	if fl.IsIgnore {
		return key, value, err
	}
	// 校验是否必填
	if !fl.IsMust {
		// 校验是否空值校验
		if fl.IsNull {
			// 根据字段类型处理
			if reflect.ValueOf(value).IsZero() {
				return key, value, errors.New(fmt.Sprintf("Err: %v -> not nil", fl.Key))
			}
		}
	}

	// 处理字段key 重置
	if fl.RenameKey != "" {
		key = fmt.Sprintf("%s", fl.RenameKey)
	} else {
		key = fmt.Sprintf("%s", fl.Key)
	}

	return key, value, err
}

func Interface2ArrayStr(obj interface{}, IsInt bool) (string, error) {
	var vs []interface{}
	buf, _ := json.Marshal(obj)
	err := json.Unmarshal(buf, &vs)
	if err != nil {
		return "", err
	}
	var value string
	if len(vs) > 0 {
		value += "("
		flen := len(vs)
		for i := 0; i < len(vs); i++ {
			if i != 0 && i != flen {
				value += ", "
			}
			if IsInt {
				value += fmt.Sprintf("%v", vs[i])
			} else {
				value += fmt.Sprintf("'%v'", vs[i])
			}
		}
		value += ")"
	} else {
		return "", errors.New("array not nil")
	}

	return value, err
}
