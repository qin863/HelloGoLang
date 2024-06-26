package main

// string

// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte uint8的别称

// rune int32的别名 表示一个Unicode 码位

//float32 float64

//complex64 complex128

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("类型：%T 值：%v\n", ToBe, ToBe)
	fmt.Printf("类型：%T 值：%v\n", MaxInt, MaxInt)
	fmt.Printf("类型：%T 值：%v\n", z, z)
}
