package selection

func Selection(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		// 当前下标
		min := i
		// 获取后续 比该值小的值的下标
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		// 判断 是否存在 后续比当前下标小的值，存在即交换位置
		if min != i {
			nums[min], nums[i] = nums[i], nums[min]
		}
	}
}
