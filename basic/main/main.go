package main

import "fmt"

func main() {
	// 短变量声明
	// 只能在函数内使用
	name := "dark"
	fmt.Print(name)

	basicType()

	defaultValue()
}

/*
go的基本类型:
	bool
	string
	int int8 int16 int32 int64
		int8 : 一个8位的int值 (-127) - 127
		int16: 应该就是16位的int值，以此类推
	uint unit8 uint16 unit32 unit64 unitptr
		uint8: 0-127
	byte // uint8 的别名
	rune // int32 的别名
		 // 表示一个Unicode码点
	float32 float64
	complex64 complex128

	int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
*/
func basicType() {
	var i = 10086
	var i8 int8 = -127
	var i16 int16 = 10086
	var ui8 uint8 = 127
	var f32 float32 = 1.10086
	// default float64? or base system?
	var f64 float64 = 1.1001010000
	fmt.Printf("i:%v i8:%v i16:%v ui8:%v \n", i, i8, i16, ui8)
	fmt.Printf("f32:%v f64:%v \n", f32, f64)
}

func defaultValue() {
	// int default 0
	// bool default false
	// string ""
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v \n", i, f, b, s)
	// print: 0 0 false (显示不出来空的字符串)
}
