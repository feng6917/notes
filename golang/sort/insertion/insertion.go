package insertion

func Insertion(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		// 未排序值
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			// 大于 未排序值 后移
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 插入未排序值
		nums[j+1] = value
	}
}
