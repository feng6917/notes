package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
值传递 数据拷贝，地址发生改变
变量重新赋值 地址不同
指针传递 地址不变

*/

func main() {
	// fmt.Println(runtime.GOARCH)

	// a := 234
	// b := 234
	// fmt.Println(&a, &b)

	// 值传递
	// var v int
	// v = 25
	// fmt.Printf("vInt: %v addr: %v size: %v type: %v\r\n", v, &v, unsafe.Sizeof(v), reflect.TypeOf(v))
	// v1 := value(v)
	// fmt.Printf("vInt: %v addr: %v size: %v type: %v\r\n", v1, &v1, unsafe.Sizeof(v1), reflect.TypeOf(v1))

	// var vuint uint
	// vuint = 255
	// fmt.Printf("vUInt: %v addr: %v size: %v type: %v\r\n", vuint, &vuint, unsafe.Sizeof(vuint), reflect.TypeOf(vuint))
	// vuint1 := valueUInt(vuint)
	// fmt.Printf("vUInt: %v addr: %v size: %v type: %v\r\n", vuint1, &vuint1, unsafe.Sizeof(vuint1), reflect.TypeOf(vuint1))

	// var vi8 int8
	// fmt.Printf("vInt8: %v addr: %v size: %v type: %v\r\n", vi8, &vi8, unsafe.Sizeof(vi8), reflect.TypeOf(vi8))
	// vi81 := valueInt8(vi8)
	// fmt.Printf("vInt8: %v addr: %v size: %v type: %v\r\n", vi81, &vi81, unsafe.Sizeof(vi81), reflect.TypeOf(vi81))

	// var vi int32
	// vi = 5
	// fmt.Printf("vInt32: %v addr: %v size: %v type: %v\r\n", vi, &vi, unsafe.Sizeof(vi), reflect.TypeOf(vi))
	// vi1 := valueInt32(vi)
	// fmt.Printf("vInt32: %v addr: %v size: %v type: %v\r\n", vi1, &vi1, unsafe.Sizeof(vi1), reflect.TypeOf(vi1))

	// var vi64 int64
	// vi64 = 255
	// fmt.Printf("vInt64: %v addr: %v size: %v type: %v\r\n", vi64, &vi64, unsafe.Sizeof(vi64), reflect.TypeOf(vi64))

	// vi641 := valueInt64(vi64)
	// fmt.Printf("vInt64: %v addr: %v size: %v type: %v\r\n", vi641, &vi641, unsafe.Sizeof(vi641), reflect.TypeOf(vi641))

	// var vuInt64 uint64
	// vuInt64 = 10
	// fmt.Printf("vuInt64: %v addr: %v size: %v type: %v\r\n", vuInt64, &vuInt64, unsafe.Sizeof(vuInt64), reflect.TypeOf(vuInt64))
	// vuInt641 := valueUInt64(vuInt64)
	// fmt.Printf("vuInt64: %v addr: %v size: %v type: %v\r\n", vuInt641, &vuInt641, unsafe.Sizeof(vuInt641), reflect.TypeOf(vuInt641))

	// var vs string
	// vs = "h"
	// fmt.Printf("value: %v addr: %v size: %v type: %v\r\n", vs, &vs, unsafe.Sizeof(vs), reflect.TypeOf(vs))
	// vs1 := valueString(vs)
	// fmt.Printf("value: %v addr: %v size: %v type: %v\r\n", vs1, &vs1, unsafe.Sizeof(vs1), reflect.TypeOf(vs1))
	arr := new([5]int)
	arr[0] = 1
	fmt.Printf("valueArray: %v addr:%v size:%v type: %v\r\n", arr, &arr, unsafe.Sizeof(arr), reflect.TypeOf(arr))
	valueArray(arr)
	fmt.Printf("valueArray: %v addr:%v size:%v type: %v\r\n", arr, &arr, unsafe.Sizeof(arr), reflect.TypeOf(arr))

	sli := make([]int, 3, 10)
	sli[0] = 1
	fmt.Printf("valueSlice: %v addr:%v size:%v type: %v\r\n", sli, &sli[0], unsafe.Sizeof(sli), reflect.TypeOf(sli))
	valueSlice(sli)
	fmt.Printf("valueSlice: %v addr:%v size:%v type: %v\r\n", sli, &sli[0], unsafe.Sizeof(sli), reflect.TypeOf(sli))

	s2 := sli[:4]
	fmt.Println(s2)
}

func valueSlice(v []int) {
	v[1] = 2
	v[2] = 3
	v = append(v, 4)
	v = append(v, 5)
	fmt.Printf("func valueSlice: %v addr:%v size:%v type: %v\r\n", v, &v[0], unsafe.Sizeof(v), reflect.TypeOf(v))
}

func valueArray(v *[5]int) {
	v[1] = 2
	v[2] = 3
	v[3] = 4
	v[4] = 5
}

func value(v int) int {
	fmt.Printf("vInt: %v addr: %v size: %v type: %v\r\n", v, &v, unsafe.Sizeof(v), reflect.TypeOf(v))
	return v + 1
}

func valueUInt(v uint) uint {
	return v + 1
}

func valueString(v string) string {
	return v
}

func valueInt8(v int8) int8 {
	return v + 1
}

func valueInt32(v int32) int32 {
	return v + 1
}

func valueInt64(v int64) int64 {
	return v + 1
}

func valueUInt64(v uint64) uint64 {
	return v + 1
}
