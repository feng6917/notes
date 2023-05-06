package bubble

func Bubble(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		ok := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				ok = true
			}
		}
		if !ok {
			break
		}
	}
}
