package main

import (
	"fmt"
	"strings"
)

func main() {
	// 切片的切片wtf?
	// 这不是二维数组吗哈哈哈
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0;i<len(board);i++ {
		fmt.Printf("%s\n",strings.Join(board[i]," "))
	}
}
