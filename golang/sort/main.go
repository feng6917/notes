package main

import (
	"fmt"
	"ssort/bubble"
	"ssort/insertion"
	"ssort/selection"
)

func main() {
	mBubble()
}

/*
冒泡排序（Bubble Sort），是一种计算机科学领域的较简单的排序算法。
它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。
走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢"浮"到数列的顶端（升序或降序排列），
就如同碳酸饮料中二氧化碳的气泡最终会上浮到顶端一样，故名"冒泡排序"。
*/
func mBubble() {
	nums := []int{43, 1, 32, 45, 56, 78}
	bubble.Bubble(nums)
	fmt.Println(nums)
}

/*
插入排序是指在待排序的元素中，
假设前面n-1(其中n>=2)个数已经是排好顺序的，现将第n个数插到前面已经排好的序列中，然后找到合适自己的位置，
使得插入第n个数的这个序列也是排好顺序的。 按照此法对所有元素进行插入，直到整个序列排为有序的过程，称为插入排序 。
*/

func mInsertion() {
	nums := []int{43, 1, 32, 45, 56, 78}
	insertion.Insertion(nums)
	fmt.Println(nums)
}

/*
选择排序法是在要排序的一组数中，选出最小（或最大）的一个数与第一个位置的数交换；
在剩下的数当中找最小的与第二个位置的数交换，即顺序放在已排好序的数列的最后，如此循环，直到全部数据元素排完为止。
*/
func mSelection() {
	nums := []int{43, 1, 32, 45, 56, 78}
	selection.Selection(nums)
	fmt.Println(nums)
}
