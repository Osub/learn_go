package main
import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T)  {
	got := Perimeter(10.0, 20.0)
	want := 60.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) //%.2f格式化浮点数保留两位小数
	} else {
		fmt.Println("测试通过矩形的周长为", got)
	}
}