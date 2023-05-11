package function

import (
	"fmt"
	"strings"
	"testing"
)

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}

// 具名函数
func Add(a, b int) int {
	return a + b
}

// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// 给函数的返回值命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

func A1() {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func() {
			println(i)
		}()

	}
}

func A2() {
	for i := 0; i < 3; i++ { // 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) {
			println(i)
		}(i)

	}
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

type Stringer interface {
	String() string
}

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

type repo struct {
	optionFirst string
}

// options options
type options func(*repo)

func NewRepo(opts ...options) *repo {
	srv := &repo{}
	for _, o := range opts {
		o(srv)
	}
	return srv
}

func withOptionFirst(t string) options {
	return func(c *repo) {
		c.optionFirst = t
	}
}
