package mysql_func

import (
	"fmt"
	"testing"
)

func TestGenDataId(t *testing.T) {
	var P struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  string `json:"age"`
	}

	_, err := GenDataId(P)
	fmt.Println(err)

	var Q struct {
		Name string `json:"name"`
		Age  string `json:"age"`
		Id   string `json:"id"`
	}

	_, err = GenDataId(Q)
	fmt.Println(err)
}

func TestUnmarshalSingleData(t *testing.T) {
	var P struct {
		Id   string `json:"id" ig:"1"`
		Name string `json:"name" im:"1" ft:"1" kv:"like" rk:"rename"`
		Age  string `json:"age" io:"1"`
	}

	_, _, _ = UnmarshalSingleData(P)
}
