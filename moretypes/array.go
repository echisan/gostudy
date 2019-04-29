package main

import "fmt"

func main() {
	var a [10]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("=============")
	slices()
}

// 切片
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// 以下操作会选择一个半开区间
	// 在数学中会以如下运用 [x,y)
	// 会返回一个数组引用
	// 切片不会存储任何数据
	var s = primes[1:4]
	fmt.Println(s)
	s[0] = 233
	fmt.Println(s)
	fmt.Println(primes)
}
