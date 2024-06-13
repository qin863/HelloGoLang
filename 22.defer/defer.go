package main

import "fmt"

// defer 语句会将函数推迟到外层函数返回之后执行

func main() {
	defer println("world")
	fmt.Println("Hello")
}
