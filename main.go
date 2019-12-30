package main
import (
	"fmt"
)

func hello(str string)  string {
	var r string = str
	return r
}

func main() {
	out := hello("testing")
	fmt.Printf("输出变量out= %s\n输出变量out指针%p \n",  out, &out)
}