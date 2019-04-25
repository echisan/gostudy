package main

import "fmt"

func main() {
	// 短变量声明
	// 只能在函数内使用
	name := "dark"
	fmt.Print(name)

	defaultVaule()
}

/*
go的基本类型:
	bool
	string
	int int8 int16 int32 int64
	uint unit8 uint16 unit32 unit64 unitptr
	byte // uint8 的别名
	rune // int32 的别名
		 // 表示一个Unicode码点
	float float64
	complex64 complex128

	int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
*/

func defaultVaule() {
	// int default 0
	// bool default false
	// string ""
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v", i, f, b, s)
	// print: 0 0 false (显示不出来空的字符串)
}
