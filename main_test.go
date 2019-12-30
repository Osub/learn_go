package mian
import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := hello()
	if want != "hello GoLang!!!" {
		t.Errorf("got %q want %q", "hello GoLang!!!", want)
	} else {
		fmt.Println("测试通过O(∩_∩)O哈哈~")
	}
}

