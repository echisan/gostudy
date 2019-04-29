package main

import "fmt"

type Local struct {
	X int
	Y int
}

func main() {
	l := Local{1,2}
	p := &l

	// 隐式引用，正常的样子应该是这个样子(*p).Y
	p.Y = 4
	(*p).X = 2

	fmt.Println(l)
}
