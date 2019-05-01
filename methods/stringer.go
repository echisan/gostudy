package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age int
}

// 其实这个是不是就是类似于java里面的toString()方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}

func main() {
	a := Person{"Arthur Dent",42}
	z := Person{"Zaphod Beeblebrox",9001}
	fmt.Println(a,z)

	b:=byte(1)
	itoa:= strconv.Itoa(int(b))
	fmt.Println(itoa)

	sprintf:= fmt.Sprintf("test %d", 1)
	fmt.Println(sprintf)
}
