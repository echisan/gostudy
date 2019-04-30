package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 函数值
	// 方法的入参可以是一个方法
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(deep(van, "deep", "dark"))
}

func van(first, second string) string {
	return first + " "+second
}

func deep(fn func(s1, s2 string) string, str1, str2 string) string {
	return fn(str1, str2)
}
