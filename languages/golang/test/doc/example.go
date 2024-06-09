package main

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
)

func main() {
	// test 变量声明
	// testVar()
	// test range
	// testRange()
	testRangeSlice()
	// 测试数组
	// testArrayEdit()
	// // 测试切片
	// testSliceEdit()

	// // 测试字符串数量
	// testStringCount()

	// // 测试for退出
	// testForExit()
	// // 测试未初始化chan插入读取
	// testNotInitChanRW()

	// // 测试结构体值修改
	// testStructEdit()

	// // 测试for值异常
	// testForValErr()

	// // test map noInit
	// // testMapNoInit()

	// // test map init
	// testMapInit()

	// // test defer exec
	// testDeferExec()
}

func testDeferExec() {
	i := [2]int{1, 2}
	fmt.Println("before i: ", i)
	defer func() {
		fmt.Println("defer print i: ", i)
	}()
	i[0] = 2
	fmt.Println("after i: ", i)
}

func testMapInit() {
	m := make(map[string]int, 1)
	m["1"] = 1
	fmt.Println(m)
	m["2"] = 2
	fmt.Println(m)
}

func testMapNoInit() {
	var m map[string]int
	m["2"] = 1 // err
	fmt.Println(m)
}

func testRange() {
	lm := make(map[int]string)
	lm[1] = "1"
	lm[2] = "2"
	lm[3] = "3"
	lm[4] = "4"
	for k, v := range lm {
		fmt.Printf("k: %d k addr: %X v:%s v addr: %X \n", k, &k, v, &v)
	}
}

func testRangeSlice() {
	/*
		range 时kv地址始终是不变的，元素的副本，引用地址发生变化
	*/
	l := []int{1, 2, 3, 4}
	for k, v := range l {
		fmt.Printf("k: %d k addr: %X v:%d v addr: %X v Addr: %X \n", k, &k, v, &v, &l[k])
	}
}

func testVar() {
	var a int
	fmt.Println(a)
	// var a, c int // err
	// fmt.Println(a, c)
	s := 1
	fmt.Println(s)
	s, sn := 2, 3
	fmt.Println(s, sn)
}

func testForValErr() {
	tsList := []TS{{"1"}, {"2"}, {Name: "3"}}
	for _, v := range tsList {
		// source
		go func() {
			v.print()
		}()
		// ans1
		// vcopy := v
		// go func() {
		// 	vcopy.print()
		// }()
		// ans2
		// go func(req TS) {
		// 	req.print()
		// }(v)
	}
	time.Sleep(time.Second)
}

type TS struct {
	Name string
}

func (c *TS) EditName(name string) {
	c.Name = name
}

func (c *TS) print() {
	fmt.Println(c.Name)
}

func testStructEdit() {
	ts := TS{Name: "1"}
	logrus.Info("edit before: ", ts)
	ts.EditName("hahah")
	logrus.Info("edit after: ", ts)
	time.Sleep(time.Second)
}

func testNotInitChanRW() {
	// var ch chan int
	// // ch <- 1
	// <-ch
	// time.Sleep(time.Second * 2)
}

func testForExit() {
	ch := make(chan struct{}, 1)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(i)
		if i == 99 {
			ch <- struct{}{}
		}
	}
	<-ch
}

func testStringCount() {
	s := "23k中国"
	logrus.Info("字符串字节数量 : ", len(s))
	logrus.Info("字符串 rune数量: ", utf8.RuneCountInString(s))

	for _, v := range []byte(s) {
		fmt.Printf("v: %#x\n", v)
	}

}

func testArrayEdit() {
	s := [5]int{1, 2, 3}
	logrus.Info("edit before: ", s)
	arrayEdit(s)
	logrus.Info("edit after: ", s)
}

func arrayEdit(s [5]int) {
	s[0] = 9
}

func testSliceEdit() {
	s := []int{1, 2, 3}
	logrus.Info("edit before: ", s)
	sliceEdit(s)
	logrus.Info("edit after: ", s)
}

func sliceEdit(s []int) {
	s[0] = 9
}
