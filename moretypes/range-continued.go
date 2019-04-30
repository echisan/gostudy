package main

import "fmt"

// 可以将下标或者该下标的值忽略掉
// 用`_`符号忽略掉你想忽略掉的值
// 如果只需要索引，直接不写第二个变量即可
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2^i
	}
	for _, v := range pow {
		fmt.Printf("%d\t", v)
	}
}
