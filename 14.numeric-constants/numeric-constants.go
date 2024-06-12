package main

import "fmt"

const (
	// 将·1 左移 100位 来创建一个非常大的数字
	// 即这个数的二进制是1 后面跟着 100个0
	Big = 1 << 100
	// 再往右移99位，即Small = 1 << 1,或者说Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1

}

func needFloat(x float64) float64 {
	return x * 0.1

}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
