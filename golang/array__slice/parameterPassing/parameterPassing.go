package parameterpassing

type Repo struct {
	Value int
}

func NewRepo(v int) *Repo {
	return &Repo{Value: v}
}

// PPArray 数组传参 值传递 函数内修改无效
func (c *Repo) PPArray(nums [1]int) {
	nums[0] = c.Value
}

// PPSlice slice传参 引用传递 函数内修改有效
func (c *Repo) PPSlice(nums []int) {
	nums[0] = c.Value
}

// PPArrayPointer 数组传参 指针传递 函数内修改有效
func (c *Repo) PPArrayPointer(nums *[1]int) {
	nums[0] = c.Value
}
