package example

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func FindPhoneNumber(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

type Repo struct {
	Value int
}

func NewRepo(v int) *Repo {
	return &Repo{Value: v}
}

// Array 数组传参 值传递 函数内修改无效
func (c *Repo) Array(nums [1]int) {
	nums[0] = c.Value
}

// Slice slice传参 引用传递 函数内修改有效
func (c *Repo) Slice(nums []int) {
	nums[0] = c.Value
}

// ArrayPointer 数组传参 指针传递 函数内修改有效
func (c *Repo) ArrayPointer(nums *[1]int) {
	nums[0] = c.Value
}

// 数组比较
func Uint64SliceEqualBCE(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// 数组去重并排序
func RemoveDuplicationMap(arr []uint64) []uint64 {
	set := make(map[uint64]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	arr = arr[:j]
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

// 数组转字符串
func uint2string(elems []uint64, joinStr string) string {
	var res []string
	switch len(elems) {
	case 0:
		return ""
	default:
		for i := range elems {
			res = append(res, strconv.FormatUint(elems[i], 10))
		}
	}
	return joinStr + strings.Join(res, joinStr)
}
