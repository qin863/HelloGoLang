package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45
	count := 1

	//while
	for count < 1000 {
		count += count

	}
	fmt.Println(count)

	count2 := 1
	for count2 < 1000 {
		count2 += count2

	}
	fmt.Println(count2)

	// 无限循环
	//for {
	//
	//}
}
