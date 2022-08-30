package main

import (
	"fmt"
)

func GenerateBoard(queens []int, n int) []string {
	board := []string{}  //board创建了个string切片
	fmt.Println(board)
	//
	for i := 0; i < n; i++ {
		row := make([]byte, n)//row创建了int切片
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))

	}
	return board
}



func main() {
	queens := make([]int, 8)
	for i := 0; i < 8; i++ {
		queens[i] = -1
	}
	fmt.Println(queens)
}
