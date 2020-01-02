package main
import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	want := Add(1, 2)
	if 3 != want {
		t.Errorf("%q测试失败%q", 3, want)
	}

	fmt.Println("测试成功")
}