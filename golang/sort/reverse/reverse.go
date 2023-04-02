package reverse

func reverse(v []interface{}) []interface{} {
	if len(v) <= 0 {
		return v
	}
	count := len(v)
	for i, j := 0, count-1; i < count/2; i++ {
		j--
		v[i], v[j] = v[j], v[i]
	}
	return v
}

// ReverseString 反转字符串
func ReverseString(v string) string {
	r := []rune(v)
	if len(r) <= 0 {
		return v
	}
	count := len(r)
	for i, j := 0, count-1; i < count/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
