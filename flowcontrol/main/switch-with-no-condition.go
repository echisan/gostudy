package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 没有条件的switch true
	// 主要能将一长串的if-then-else写的更加清晰
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")

	}
}

