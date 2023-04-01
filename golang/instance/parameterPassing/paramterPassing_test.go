package parameterpassing


import (
	"testing"
)


func TestNewParameter(t *testing.T){
	srv := NewParameter(ParameterOne("haha"))
	got := srv.ParameterOne
	want := "haha"
	if got != want{
		t.Errorf("want %s but got %s", want, got)
	}
}