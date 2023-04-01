package util

func ErrorStr(v string) string {
	beforeStr := "\r\n-----------------------------------------------\r\n"
	afterStr := "\r\n-----------------------------------------------\r\n"
	return beforeStr + "执行报错: " + v + "\r\n请重试!!!" + afterStr
}
