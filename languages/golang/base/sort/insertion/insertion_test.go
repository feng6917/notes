package insertion

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	nums := []int{43, 1, 32, 45, 56, 78}
	Insertion(nums)
	fmt.Println(nums)
}
