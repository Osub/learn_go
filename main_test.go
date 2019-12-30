package main
import (
	"fmt"		//引入格式化包
	"testing" //引入测试包
)

func TestHello(t *testing.T) {
	var str string
	str = "hello GoLang!!!"
	want := hello(str)
	if want != str {
		t.Errorf("got %q want %q", "hello GoLang!!!", want)
	} else {
		fmt.Println("测试通过O(∩_∩)O哈哈~")
	}
}

