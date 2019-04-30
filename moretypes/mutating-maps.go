package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]string)
	// 插入元素
	m["code"] = "1"

	elem := m["code"]

	fmt.Println("code:" + elem)
	// 删除元素
	delete(m,"code")
	fmt.Println(m["code"])
	elem,p:=m["code"]
	fmt.Println(p)
}
