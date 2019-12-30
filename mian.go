package mian
import (
	"fmt"
)

func hello()  string {
	return "hello GoLang!!!"
}

func main() {
	out := hello()
	fmt.Printf("输出变量out= %s\n输出变量out指针%p ",  out, &out)
}