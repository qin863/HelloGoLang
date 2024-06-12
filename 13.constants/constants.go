package main

import "fmt"

const Pi = 3.14

// 常量 使用const
// 不能使用 := 语法声明

func main() {
	const world = "世界"
	fmt.Println("Hello", world)     //Hello 世界
	fmt.Println("Happy", Pi, "Day") // Happy 3.14 Day

	const Truth = true
	fmt.Println("GO rules?", Truth) //GO rules? true
}
