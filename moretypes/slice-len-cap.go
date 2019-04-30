package main

import "fmt"

func main() {
	// 切片拥有长度(len)和容量(cap(capacity))
	// 感觉有点像java-nio中的limit跟capacity

	// 定义一个切片
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len:%v cap:%v slice:%v\n", len(s), cap(s), s)

	s = s[0:5]
	fmt.Printf("len:%v cap:%v\n slice:%v\n", len(s), cap(s), s)

	// 给切片扩展长度
	s = s[0:6]
	fmt.Printf("len:%v cap:%v\n slice:%v\n", len(s), cap(s), s)

	// 舍去前面两个
	s = s[2:]
	fmt.Printf("len:%v cap:%v\n slice:%v\n", len(s), cap(s), s)

}
