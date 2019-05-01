package main

import (
	"fmt"
	"math"
)

type Vertey struct {
	X, Y float64
}

func (y Vertey) Abs() float64 {
	return math.Sqrt(y.X*y.X + y.Y*y.Y)
}

func (y *Vertey) Scale(f float64) {
	y.X = y.X * f
	y.Y = y.Y * f
}

func main() {
	v := Vertey{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())

	a := 1
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

}



