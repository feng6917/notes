package mysql_func

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"reflect"
)

/*
is_ignore    -> ig 值是否忽略
is_must      -> im 值是否必填
is_null      -> in 值是否空值校验
field_type   -> ft 字段类型
key_value    -> kv 字段关系
rename_key   -> rk 字段key重命名
is_or        -> io 字段间关系

*/

const (
	FieldTypeString      = "1" // string
	FieldTypeInt         = "2" // int8 int32 int64 uint32 uint64
	FieldTypeBool        = "3" // bool
	FieldTypeIntArray    = "4" // []int8 []int32 []int64 []uint64
	FieldTypeStringArray = "5" // []string
)

const (
	FieldKeyValueEqual     = "="
	FieldKeyValueNotEqual  = "<>"
	FieldKeyValueMoreThan  = ">"
	FieldKeyValueMoreEqual = ">="
	FieldKeyValueLessThan  = "<"
	FieldKeyValueLessEqual = "<="
	FieldKeyValueIn        = "in"
	FieldKeyValueNotIn     = "not in"
	FieldKeyValueLike      = "like"
)

var GlobalDC = make(map[string][]*FieldData, 0)

type FieldData struct {
	Index     int    `json:"index"`      // 索引
	Key       string `json:"key"`        // key
	Value     string `json:"value"`      // value
	IsIgnore  bool   `json:"is_ignore"`  // 值是否忽略
	IsMust    bool   `json:"is_must"`    // 值是否必填
	IsNull    bool   `json:"is_null"`    // 值是否空值校验 string != ""; int != 0; [] > 0; time != 0;
	FieldType string `json:"field_type"` // 值类型
	KeyValue  string `json:"key_value"`  // 对应关系 = < > <= >= like in not in
	RenameKey string `json:"rename_key"` // key 重置
	Sql       string `json:"sql"`        // 处理后sql语句
	IsOr      bool   `json:"is_or"`      // 查询状态是否是or关系
}

// 初始化解析数据
func UnmarshalData(data []interface{}, searchStr ...string) error {
	if len(data) > 0 {
		for _, sd := range data {
			id, fieldList, err := UnmarshalSingleData(sd, searchStr...)
			if err != nil {
				return err
			}
			GlobalDC[id] = fieldList
		}
	}
	return nil
}

//  解析单条数据
func UnmarshalSingleData(obj interface{}, searchStr ...string) (string, []*FieldData, error) {
	// searchStr 对应值为数据库字段;默认searchStr为json
	searchStrValue := "json"
	if len(searchStr) > 0 {
		searchStrValue = searchStr[0]
	}
	var err error
	var id string
	var fieldList []*FieldData
	id, err = GenDataId(obj)
	if err != nil {
		return id, fieldList, err
	}
	// 解析单条数据 tag
	st := reflect.TypeOf(obj)
	fieldNum := st.NumField()
	for i := 0; i < fieldNum; i++ {
		filed := st.Field(i)
		f := FieldData{}
		f.Index = i
		// 处理值是否忽略
		tig := filed.Tag.Get("ig")
		if tig != "" {
			f.IsIgnore = true
		} else {
			// 处理key 字段必填
			f.Key = filed.Tag.Get(searchStrValue)
			// 处理值是否必填
			tim := filed.Tag.Get("im")
			if tim != "" {
				f.IsMust = true
			} else {
				// 处理值是否空值校验
				tin := filed.Tag.Get("in")
				if tin != "" {
					f.IsNull = true
				}
			}
			// 处理kv 对应关系 默认 =
			tkv := filed.Tag.Get("kv")
			if tkv != "" {
				f.KeyValue = tkv
			} else {
				f.KeyValue = "="
			}
			// 处理字段key 重置
			trk := filed.Tag.Get("rk")
			if trk != "" {
				f.RenameKey = trk
			}
			// 处理字段类型
			tft := filed.Tag.Get("ft")
			if tft != "" {
				f.FieldType = tft
			} else {
				f.FieldType = FieldTypeString
			}
			// 处理查询状态对应关系
			tio := filed.Tag.Get("io")
			if tio != "" {
				f.IsOr = true
			}
		}
		fieldList = append(fieldList, &f)
	}
	return id, fieldList, err
}

func GenDataId(obj interface{}) (string, error) {
	var tagStr string
	st := reflect.TypeOf(obj)
	fieldNum := st.NumField()
	for i := 0; i < fieldNum; i++ {
		filed := st.Field(i)
		tagStr += filed.Tag.Get("json")
	}
	hash := md5.New()
	for buf, reader := make([]byte, 65536), bufio.NewReader(bytes.NewReader([]byte(tagStr))); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
