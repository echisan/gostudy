package main

import "fmt"

func main() {
	// 这是一个数组
	arr := [3]int{1, 2, 3}
	// 这是一个切片
	// 不定义数组大小就是切片？
	sli := []bool{true, false, true}

	fmt.Println(arr)
	fmt.Println(sli)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}

	fmt.Println(s)
}
