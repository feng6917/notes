package mysql_func

type MysqlCache struct {
	List []Data
}

type Data struct {
	Name string
	Map  map[string]interface{}
}

//func UnmarshalInterface(i interface{}) map[string]interface{} {
//	reflect.ValueOf()
//}
