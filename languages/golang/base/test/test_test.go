package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("全局测试前执行操作,初始化服务、数据库及创建数据等")
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()
	fmt.Println("全局测试前执行操作,关闭服务、数据库及删除数据等")
	os.Exit(retCode)
}

// 开始及结束操作
func setupTestSum(t *testing.T) func(t *testing.T) {
	t.Log("setup start")
	return func(t *testing.T) {
		t.Log("setup end")
	}
}

// 单个
func TestSum(t *testing.T) {
	n1, n2 := 1, 1
	got := Sum(n1, n2)
	want := 2
	setupTestSum(t)
	defer setupTestSum(t)
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

// 多个（组）
func TestSum2(t *testing.T) {
	type test struct {
		inputN1 int
		inputN2 int
		want    int
	}
	tests := map[string]test{
		"true":  {1, 1, 2},
		"false": {1, 1, 3},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Sum(tc.inputN1, tc.inputN2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 基准测试并不会默认执行，需要增加-bench参数 go test -bench=Sum
func BenchmarkSum(b *testing.B) {
	b.ResetTimer() // 重置时间
	for i := 0; i < b.N; i++ {
		Sum(1, 1)
	}
}

// 以并行的方式执行给定的基准测试。
func BenchmarkSumParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sum(1, 1)
		}
	})
}

// 示例函数
func ExampleSum() {
	fmt.Println(Sum(1, 2))
	// Output:
	// 3
}
