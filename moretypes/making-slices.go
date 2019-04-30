package main

import "fmt"

func main() {
	// 切片可以用make来创建，这是创建动态数组的方式
	a := make([]int, 5) // 这tm动态数组？
	fmt.Printf("len:%v cap:%v slice:%v\n", len(a), cap(a), a)

	// 如果要指定容量
	// 第二个参数是长度，第三个是容量
	b := make([]int,0,5)
	fmt.Printf("len:%v cap:%v slice:%v\n", len(b), cap(b), b)

	b = b[:5]
	fmt.Printf("len:%v cap:%v slice:%v\n", len(b), cap(b), b)


}
