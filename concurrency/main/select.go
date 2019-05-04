package main

import (
	"fmt"
	"time"
)

/*
select语句使一个Go程可以等待多个通信操作
select会阻塞到某个分支可以继续执行为止

go还能发现死锁的
fatal error: all goroutines are asleep - deadlock!

如果不符合当中所有的case则会一直执行default中定义好的方法
 */
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for true {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("   zzz.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	fibonacci2(c, quit)
}
