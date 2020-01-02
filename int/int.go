package main

import (
	"fmt"
)

func Add(a int, b int) int {
	c := a + b
	return c
}

func main()  {
	r := Add(1, 2)
	fmt.Println("a+b的计算结果是", r)
}