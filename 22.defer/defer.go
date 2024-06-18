package main

import "fmt"

// defer 语句会将函数推迟到外层函数返回之后执行

func main() {
	var s = "world"
	defer println(s)
	s = "1"
	fmt.Println("Hello")
}
