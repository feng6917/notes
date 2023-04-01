package excelTag

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

const (
	ZeroInt = 0
	TwoInt  = 2
)

type WriteXlsx struct {
	Buffer    bytes.Buffer
	SheetName string
	Obj       interface{}
	Funcs     map[string]interface{}
}

func (c *WriteXlsx) WriteXlsxFunc() error {
	// 解析结构
	sm := unmarshalStructType(c.Obj)
	// 生成新的xlsx文件
	excelFile := excelize.NewFile()
	if c.SheetName != "" {
		excelFile.NewSheet(c.SheetName)
		excelFile.DeleteSheet("Sheet1")
	} else {
		c.SheetName = "Sheet1"
	}
	var err error
	// 写入header
	for _, v := range sm {
		err = excelFile.SetCellStr(c.SheetName, fmt.Sprintf("%s%d", v.ColV, 1), v.RnV)
		if err != nil {
			return err
		}
	}
	// 写入数据
	sv := c.unmarshalStructValue(sm, c.Obj)
	if len(sv) > 0 {
		for _, v := range sv {
			err = excelFile.SetCellStr(c.SheetName, v.Axis, v.NV)
			if err != nil {
				return err
			}
		}
	}

	_, err = excelFile.WriteTo(bufio.NewWriter(&c.Buffer))
	if err != nil {
		return err
	}
	return nil
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func (c *WriteXlsx) Call(funcName string, param string) (string, error) {
	var val string
	if c.Funcs[funcName] == nil {
		return "", errors.New("func not exist")
	}
	fn := reflect.ValueOf(c.Funcs[funcName])
	result := fn.Call([]reflect.Value{reflect.ValueOf(param)})
	if len(result) == 1 {
		val = result[0].String()
	}
	return val, nil
}

type StructType struct {
	Index int               // 下标
	JsV   string            // json 值
	RnV   string            // reName 值
	FnV   string            // func 值
	EnV   map[string]string // enum map 0|保密;1|男;2|女
	Type  reflect.Type      // 类型
	ColV  string            // 列名
}

func unmarshalStructType(obj interface{}) map[int]StructType {
	sm := make(map[int]StructType, ZeroInt)
	el := reflect.TypeOf(obj).Elem()
	index := 0
	for i := 0; i < el.NumField(); i++ {
		tg := el.Field(i).Tag
		igv := tg.Get("ig")
		if igv != "1" {
			enm := make(map[string]string)
			env := tg.Get("en")
			if strings.Contains(env, "|") {
				if strings.Contains(env, ";") {
					ms := strings.Split(env, ";")
					for _, k := range ms {
						es := strings.Split(k, "|")
						if len(es) == TwoInt {
							enm[es[0]] = es[1]
						}
					}
				} else {
					es := strings.Split(env, "|")
					if len(es) == TwoInt {
						enm[es[0]] = es[1]
					}
				}
			}
			index += 1
			ic := toChar(index)
			sm[i] = StructType{
				Index: index,
				JsV:   tg.Get("json"),
				RnV:   tg.Get("rn"),
				FnV:   tg.Get("fn"),
				EnV:   enm,
				Type:  el.Field(i).Type,
				ColV:  strings.ReplaceAll(fmt.Sprintf("%q", ic), "'", ""),
			}
		}
	}
	return sm
}

type StructValue struct {
	Axis string // 坐标
	NV   string // 新值
}

func (c *WriteXlsx) unmarshalStructValue(sm map[int]StructType, objs interface{}) []StructValue {
	sv := make([]StructValue, 0)
	ev := reflect.ValueOf(objs)
	if ev.Len() > 0 {
		for i := 0; i < ev.Len(); i++ {
			v := ev.Index(i)
			for j := 0; j < v.NumField(); j++ {
				st := sm[j]
				if st.Index != 0 && st.ColV != "" {
					var nv interface{}
					ty := st.Type.Kind()
					switch ty {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						nv = v.Field(j).Int()
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
						nv = v.Field(j).Uint()
					case reflect.Float32, reflect.Float64:
						nv = v.Field(j).Float()
					case reflect.String:
						nv = v.Field(j).String()
					case reflect.Bool:
						nv = v.Field(j).Bool()
					default:
						nv = v.Field(j).String()
					}
					// 获取方法值
					if st.FnV != "" {
						str, err := c.Call(st.FnV, fmt.Sprintf("%v", nv))
						if err == nil {
							nv = str
						}
					}
					// 获取枚举值
					if len(st.EnV) > 0 {
						nv = st.EnV[fmt.Sprintf("%v", nv)]
					}
					sv = append(sv, StructValue{
						Axis: fmt.Sprintf("%s%d", st.ColV, i+2),
						NV:   fmt.Sprintf("%v", nv),
					})
				}
			}
		}
	}
	return sv
}
