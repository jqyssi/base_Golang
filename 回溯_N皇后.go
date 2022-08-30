package main

import "fmt"
//二维,[[board],[board]...]
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	//queens切片里的数是对应正确的下标
	queens := make([]int, n)
	//把queen切片全部置为—1
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	//三个map
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}
/**
 * @Description :(皇后切片slice  皇后个数int  行int  列map  撇map 捺map)
 * @params :
 * @Author : ranksin
 * @Date : 2022/5/23 10:19 上午
 */
func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, n, row + 1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}
/**
 * @Description :一行一行创建棋盘 刚开始一行棋盘全部初始化为"." 然后根据queen切片中的数字放置皇后位置，也就是说Queen切片为正确放置的一组数
 * @params :参数（正确棋盘列数）
 * @Author : ranksin
 * @Date : 2022/5/23 9:24 上午
 */
func generateBoard(queens []int, n int) []string {
	board := []string{}  //生成一行字符串切片

	//i表示行循环
	for i := 0; i < n; i++ {
		//每一行创建一个字符串切片
		row := make([]byte, n)
		//将每一行的每一列全部初始化为 "."
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		//将每一行的对应的queen下标设置为皇后位置
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}

	return board
}

func main() {

	fmt.Println(solveNQueens(8))
	fmt.Println("所有的解有",len(solutions),"种")
}