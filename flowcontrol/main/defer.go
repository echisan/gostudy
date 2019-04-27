package main

import "fmt"

func main() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	// 参数会立即求值，但是得等待外层函数返回前才会调用这个defer方法
	defer fmt.Println("world")
	fmt.Println("hello")
}
