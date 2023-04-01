package parameterpassing

func PPArray(nums [3]int) {
	nums[0] = 111
}

func PPSlice(nums []int) {
	if len(nums) > 0 {
		nums[0] = 111
	}
}

func PPArray2(nums *[3]int) {
	nums[0] = 111
}
