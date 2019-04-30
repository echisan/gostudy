package main

import "fmt"

// 终于到了java中了hashMap了吧
type Vertex struct {
	Lat, Long float64
}
type Person struct {
	Name, Sex string
}

// java中大概这个样子
// Map<String,Vertex> map = new HashMap<>();
var m map[string]Vertex

// 初始化一个map
var p  = map[int]Person{
	1:{"dick","boy"},
	2:{"ass","boy"},
}

func main() {
	// make就等于new了吧
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433,-74.39967}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	fmt.Println(p)
	p[3] = Person{"test","girl"}
	fmt.Println(p)

	for _,v:=range p{
		fmt.Println(v)
	}
}
