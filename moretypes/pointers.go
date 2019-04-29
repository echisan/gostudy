package main

import "fmt"

func main() {
	i := 42

	// 通过&操作符会生成一个指向其操作数的指针
	p := &i

	fmt.Printf("i: %v p: %v\n", i, p)

	// 通过*操作符表示指针指向的底层值

	pv := *p
	fmt.Printf("i:%v p:%v pv:%v\n", i, p, pv)

	// 使用*p 可以修改i的指
	*p = 21
	fmt.Printf("i:%v p:%v pv:%v", i, p, pv)

}
