package bubble

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	nums := []int{43, 1, 32, 45, 56, 78}
	Bubble(nums)
	fmt.Println(nums)
}
