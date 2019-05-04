package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

type UnsafeCounter struct {
	v map[string]int
}


func (c *UnsafeCounter) Inc(key string) {
	c.v[key]++
}

func (c *UnsafeCounter) Value(key string) int{
	return c.v[key]
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int{
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	//c:=SafeCounter{v:make(map[string]int)}
	u:=UnsafeCounter{v:make(map[string]int)}
	//for i := 0; i < 1000; i++ {
	//	go c.Inc("somekey")
	//}
	for i := 0; i < 1000; i++ {
		go u.Inc("somekey")
	}
	//time.Sleep(time.Second)
	//fmt.Println(c.Value("somekey"))

	// fatal error: concurrent map writes
	fmt.Println(u.Value("somekey"))
}
