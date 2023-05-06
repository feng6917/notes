package reverse

import "testing"

func TestReverseString(t *testing.T) {
	s := "123456"
	want := "654321"
	got := ReverseString(s)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
