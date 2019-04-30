package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "how are you"

	// 就是java重的spilt()方法
	fields:= strings.Fields(str)
	for _,s:=range fields{
		fmt.Println(s)
	}
}
