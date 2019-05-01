package main

import (
	"fmt"
	"math"
)

type User struct {
	X,Y float64
}

// 方法就是一类带特殊的**接收者**参数的函数
func (u User) Abs() float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}
// 方法即函数
func Abs(u User) float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}

func main() {
	u:=User{3,4}
	fmt.Println(u.Abs())

	// 其实是不是 u.Abs() 就等价于 Abs(u) 呢
	Abs(u)
}

