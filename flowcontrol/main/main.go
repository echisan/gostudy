package main

import "fmt"

func main() {
	forTest()
	forTest2()
	whileTest()
	//forever()
	ifTest(5)
	ifWithShortStatement(20, 30)

}

func forTest() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forTest2() {
	// 初始化语句和后置语句是可选的
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func whileTest() {
	// go 中 for 即 while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forever() {
	// 无限循环
	// java 中是while(true)
	for {
		fmt.Println("forever")
	}
}

func ifTest(num int) {
	if num < 10 {
		fmt.Println("print if true")
	}
}

func ifWithShortStatement(x, y int) {
	if z := x + y; z < 100 {
		fmt.Println(z)
	} else {
		// else 中也可使用z
		fmt.Println(x - y + z)
	}

}
