package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	var expected string = "aaaaa"
	var r string
	r = repeat()
	if r != expected {
		t.Errorf("expected %q but got %q", expected, r)
	} else {
		fmt.Println("测试通过")
	}
}

//基准测试
func BenchmarkRepeat(b *testing.B) {
	for index := 0; index < b.N; index++ {
		repeat()	
	}
}