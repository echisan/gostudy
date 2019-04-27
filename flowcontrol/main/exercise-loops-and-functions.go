package main

import (
	"fmt"
)

/*
	实现一个平方根函数
 */
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 7; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	Sqrt(16)
	fmt.Println("-------------")
	Sqrt(36)
}
