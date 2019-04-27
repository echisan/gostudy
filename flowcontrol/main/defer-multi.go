package main

import "fmt"

func main() {
	fmt.Println("counting")
	// defer栈 被推迟的函数会压入栈中
	// 当外层函数返回前再以后进先出的顺序调用
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
