package main
import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("测试切片的和", func(t *testing.T) {
		want := 9
		var r int
		r, _ = sum([]int{1, 3, 2, 2, 0, 1})
		if want != r {
			t.Errorf("got %q want %q", r, want)
		} else {
			fmt.Println("O(∩_∩)O哈哈~，测试通过了")
		}
	})
}