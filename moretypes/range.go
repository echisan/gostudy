package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

var sub = []string{"java", "python", "golang", "php", "javascript"}

// range可以遍历切片或者映射
// 第一个值为 下标 第二的值为 该下标对应元素的一份**副本**
func main() {
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	for i, v := range sub {
		fmt.Printf("%s is No.%d\n", v, i+1)
	}

}
