package main

import (
	"fmt"
	"os"
	"regexp"
)

// ps: 照搬 Go语言高级教程 仅用于笔记记录

/*
	数组是一种值类型，虽然数组的元素可以被修改，但是数组 本身的赋值和函数传参都是以整体复制的方式处理的。

	字符串底层数据也是对应的字节数组，但是字符串的只读属性禁止了在程序中对底层字节数组的元素的修改。
	字符串赋值 只是复制了数据地址和对应的长度，而不会导致底层数据的复制。

    切片的行为更为灵活，切片的结构和字符串结构类似，但是解除了只读限制。
	切片的底层数据虽然也是对应数据类型的 数组，但是每个切片还有独立的长度和容量信息，切片赋值和 函数传参数时也是将切片头信息部分按传值方式处理。
	切片头含有底层数据的指针，所以它的赋值也不会导致底层数据 的复制。

*/

func main() {
	// array
	// 数组是一个由固定长度的特定类型元素组成的序列，一个数组 可以由零个或多个元素组成。数组的长度是数组类型的组成部分。
	// 因为数组的长度是数组类型的一个部分，不同长度或不同 类型的数据组成的数组都是不同的类型，因此在Go语言中很 少直接使用数组(不同长度的数组因为类型不同无法直接赋 值)。

	// 数组定义方式:
	var a [3]int              // int 型数组, 元素全部为0
	var b = [...]int{1, 2, 3} // int类型数组, 元素为 1, 2, 3
	// 定义一个长度为3的int类 //定义一个长度为3的
	var c = [...]int{2: 3, 1: 2}    //定义一个长度为3 的int类型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} //定义一个长度为 6的int类型数组, 元素为 1, 2, 0, 0, 5, 6
	fmt.Println(a, b, c, d)

	var e = [...]int{1, 2, 3} // e是一个数组
	var f = &e                // f 是指向数组的指针
	fmt.Println(e[0], e[1])   // 打印数组的前2个元素
	fmt.Println(f[0], f[1])   // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range f { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	// string
	// 一个字符串是一个不可改变的字节序列，字符串通常是用来包含人类可读的文本数据。
	// 字符串的元素不可 修改，是一个只读的字节数组。
	// 每个字符串的长度虽然也是固 定的，但是字符串的长度并不是字符串类型的一部分。字符串可以包含任意的数据.

	// slice
	// 切片就是一种简化版的动态数组。因为动态数组的长度是不固定，切片的长度自然也就不能是类型的组成部分了。
	// 切片定义方式:
	var (
		a_ []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
		b_ = []int{}           // 空切片,和nil不相等,一般用来 表示一个空的集合
		c_ = []int{1, 2, 3}    // 有3个元素的切片,len和cap 都为3
		d_ = c[:2]             // 有2个元素的切片, len为2, cap为3
		e_ = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
		f_ = c[:0]             // 有0个元素的切片, len为0, cap为3
		g_ = make([]int, 3)    // 有3个元素的切片, len和cap都为3
		h_ = make([]int, 2, 3) // 有2个元素的切片,len为2, cap为3
		i_ = make([]int, 0, 3) // 有0个元素的切片,len为0, cap为3
	)
	fmt.Println(a_, b_, c_, d_, e_, f_, g_, h_, i_)
	fmt.Println("1")

	// 添加切片元素
	a_ = append(a_, 1)                 // 追加1个元素
	a_ = append(a_, 1, 2, 3)           //追加多个元素,手写解 包方式
	a_ = append(a_, []int{1, 2, 3}...) //追加一个切片,切片 需要解包

	a_ = append([]int{0}, a_...)          // 在开头添加1个元素
	a_ = append([]int{-3, -2, -1}, a_...) //在开头添加1个切片

	i := 2
	x := 2
	a_ = append(a_[:i], append([]int{x}, a_[i:]...)...)       // 在第i个位置插入x
	a_ = append(a_[:i], append([]int{1, 2, 3}, a_[i:]...)...) // 在第i个位置插入切片

	a_ = append(a_, 0)     // 切片扩展1个空间
	copy(a_[i+1:], a_[i:]) // a[i:]向后移动1个位置
	a_[i] = x              // 设置新添加的元素

	// 删除切片元素
	N := 2
	a_ = a_[:len(a_)-1] // 删除尾部1个元素
	a_ = a_[:len(a_)-N] // 删除尾部N个元素

	a_ = a_[1:] //删除开头1个元素
	a_ = a_[N:] // 删除开头N个元素

	a_ = append(a_[:0], a_[1:]...) // 删除开头1个元素
	a_ = append(a_[:0], a_[N:]...) // 删除开头N个元素

	a_ = append(a_[:i], a_[i+1:]...) // 删除中间1个元素
	a_ = append(a_[:i], a_[i+N:]...) // 删除中间N个元素

	a_ = a_[:i+copy(a_[i:], a_[i+1:])] // 删除中间1个元素
	a_ = a_[:i+copy(a_[i:], a_[i+N:])] // 删除中间N个元素
	fmt.Println(a_)

	// 对于切片来说， len 为 0 但是cap 容量不为 0 的切片则是非常有用的特性。
	// 切片高效操作的要点是要降低内存分配的次数，尽量保证 append 操作不会超出 cap 的容量，降低触发内存分配的次数和每次分配内存大小。
	// Filter()

	// 避免切片内存泄漏
	// FindPhoneNumber()

	// 删除切片元素时可能会遇到
	var ae []*int
	ae[len(ae)-1] = nil // GC回收最后一个元素内存
	ae = ae[:len(ae)-1]
	fmt.Println(ae)

}

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
