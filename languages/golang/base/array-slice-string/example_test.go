package example

import "testing"

func TestArray(t *testing.T) {
	r := NewRepo(28)
	nums := [1]int{0}
	r.Array(nums)
	if r.Value == nums[0] {
		t.Errorf("got %d, want %d", nums[0], 0)
	}
}

func TestRepo_Slice(t *testing.T) {
	r := NewRepo(28)
	nums := []int{0}
	r.Slice(nums)
	if r.Value != nums[0] {
		t.Errorf("got %d, want %d", nums[0], r.Value)
	}
}

func TestRepo_ArrayPointer(t *testing.T) {
	r := NewRepo(28)
	nums := [1]int{0}
	r.ArrayPointer(&nums)
	if r.Value != nums[0] {
		t.Errorf("got %d, want %d", nums[0], r.Value)
	}
}
