package main

import "fmt"

// f(n)=f(n-1)+f(n-2)
// f(0) = 0
// f(1) = 1
// f(2) = f(1) + f(0) = 1 + 0 = 1
func fibonacci() func() int {
	a,b := 0,1
	return func() int {
		// 第一次返回 0，此时a = 1,b = 0+1 = 1
		// 第二次返回 1, 此时a = 1,b = 1+1 = 2
		c:=a
		a,b = b,a+b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
