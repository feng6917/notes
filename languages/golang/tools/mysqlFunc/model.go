package mysql_func

type IRepo interface {
	// 校验非空
	CheckSpace(i interface{}, ft int) bool
	// 校验必填
	// 校验是否可忽略
	// 校验是否重命名

}

// 校验非空
//func ToDoSpace(tn string) error {
//	// 获取缓存
//	var ok bool
//	if ft == FieldTypeInt {
//		v := fmt.Sprintf("%v", i)
//		if v == "" || v == "0" {
//			ok = true
//		}
//	} else if ft == FieldTypeString {
//		v := fmt.Sprintf("%v", i)
//		if v == "" {
//			ok = true
//		}
//	}
//	return ok
//}

// 校验必填
func ToDoMust() {

}
