package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s,ok := i.(string)
	fmt.Println(s,ok)

	f,ok := i.(float64)
	fmt.Println(f,ok)

	// 如果没有ok这个参数会报错
	f = i.(float64)
	fmt.Println(f)
}
