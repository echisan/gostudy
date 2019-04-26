package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 100
	Small = Big >> 99 // small = 2
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
	// constant 1267650600228229401496703205376 overflows int
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
	fmt.Println(math.MaxInt64)
}
