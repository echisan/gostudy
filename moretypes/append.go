package main

import "fmt"

func main() {
	// 向切片追加元素
	s := make([]int, 0, 5)
	printBasic(s)

	// 当给定的切片底层数组太小，不足以容纳给定的值
	// 它会分配更大的数组
	// 返回的切片会指向这个新分配的数组
	s = append(s, 1, 2, 3, 4, 5, 6)
	printBasic(s)

	fmt.Printf("s type:%T", s)
}

func printBasic(s []int) {
	fmt.Printf("len:%v cap:%v slice:%v \n", len(s), cap(s), s)
}
