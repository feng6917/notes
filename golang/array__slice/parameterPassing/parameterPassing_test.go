package parameterpassing

import "testing"

func TestPPArray(t *testing.T) {
	r := NewRepo(28)
	nums := [1]int{0}
	r.PPArray(nums)
	if r.Value == nums[0] {
		t.Errorf("got %d, want %d", nums[0], 0)
	}
}

func TestRepo_PPSlice(t *testing.T) {
	r := NewRepo(28)
	nums := []int{0}
	r.PPSlice(nums)
	if r.Value != nums[0] {
		t.Errorf("got %d, want %d", nums[0], r.Value)
	}
}

func TestRepo_PPArrayPointer(t *testing.T) {
	r := NewRepo(28)
	nums := [1]int{0}
	r.PPArrayPointer(&nums)
	if r.Value != nums[0] {
		t.Errorf("got %d, want %d", nums[0], r.Value)
	}
}
