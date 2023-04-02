package optionFunc

import (
	"testing"
)

func TestNewRepo(t *testing.T) {
	want := "optionFirst"
	srv := NewRepo(withOptionFirst(want))
	got := srv.optionFirst
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
