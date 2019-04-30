package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	// 切片的零值为nil
	if s == nil {
		fmt.Println("nil")
	}
}
