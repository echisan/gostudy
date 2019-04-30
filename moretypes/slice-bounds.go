package main

import "fmt"

// 切片的默认行为
// 默认下界为0 上界为长度
func main() {
	var a [10]int

	s1 := a[0:10]
	s2 := a[0:]
	s3 := a[:10]
	s4 := a[:]

	fmt.Printf("s1:%v\ns2:%v\ns3:%v\ns4:%v", s1, s2, s3, s4)
}
