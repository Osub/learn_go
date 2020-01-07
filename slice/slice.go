package main
import (
	"fmt"
)

func sum(numbers []int) (int, []int) {
	var r int
	var res []int
	for _, number := range numbers {
		r += number
		res = append(res, r)
	} 
	return r, res
}

func main() {
	v := []int{3, 6, 10}
	r, res := sum(v)
	fmt.Println("计算结果是", r,"切片是",  res)
}