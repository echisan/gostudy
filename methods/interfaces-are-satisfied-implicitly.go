package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

//func main() {
//	var i I
//
//	var t *T
//	i = t
//	describe(i)
//	i.M()
//
//	var t1 T
//	fmt.Printf("t1 %v %T nil:%v",t1,t1,&t1==nil)
//
//	i = &T{"hello"}
//	describe(i)
//	i.M()
//
//	i = F(math.Pi)
//	describe(i)
//	i.M()
//}

func main() {
	var i I
	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}
