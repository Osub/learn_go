package main
import (
	"fmt"
)

func repeat() string {
	count := 5
	str := "a"
	var r string
	r = ""
	for index := 0; index < count; index++ {
		r += str 
	}
	return r
}

func strings()  {
	//strings标准包的常用函数
	
}

func main()  {
	str := repeat()
	fmt.Println("输出结果是", str)	
}