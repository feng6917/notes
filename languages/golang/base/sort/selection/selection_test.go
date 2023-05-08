package selection

import (
	"fmt"
	"testing"
)

func TestSelection(t *testing.T) {
	nums := []int{43, 1, 32, 45, 56, 78}
	Selection(nums)
	fmt.Println(nums)
}
