package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type EError struct {
	What string
}

func (e *EError) Error() string {
	return fmt.Sprintf("what %v",e.What)
}

func testError() error {
	return &EError{"the fuck"}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s",e.When,e.What)
}

func run() error {
	return &MyError{time.Now(),
		"it didn't work"}
}

func main() {
	//if err := run(); err != nil {
	//	fmt.Println(err)
	//}

	i,err:=strconv.Atoi("42i")
	if err != nil {
		fmt.Printf("couldn't convert number:%v\n",err)
	}else {
		fmt.Println("Converted integer:",i)
	}


	fmt.Println(testError())
}
