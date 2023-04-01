package main

import (
	"arrayWithSlice/parameterpassing"
	"fmt"
)


func main() {
	

}

// mParameterPassing 数组与切片传参
func mParameterPassing() {
	// 数组传参 值传递 函数内修改无效
	nums := [3]int{1, 2, 3}
	parameterpassing.PPArray(nums)
	fmt.Println("array nums: ", nums)

	// slice传参 引用传递 函数内修改有效
	nums1 := []int{1, 2, 3}
	parameterpassing.PPSlice(nums1)
	fmt.Println("slice nums: ", nums1)

	// 数组传参 指针传递 函数内修改有效
	nums2 := [3]int{1, 2, 3}
	parameterpassing.PPArray2(&nums2)
	fmt.Println("array pointer nums: ", nums2)
}