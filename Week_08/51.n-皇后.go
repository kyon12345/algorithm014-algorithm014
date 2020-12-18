/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
package main

import (
	"image/jpeg"
	"math/bits"
	"runtime/trace"
	"strings"
)

// @lc code=start
// var ans [][]string

// var cols map[int]bool
// var pie map[int]bool
// var na map[int]bool

// func solveNQueens(n int) [][]string {
// // 	if n <= 0 {
// // 		return ans
// // 	}
// // 	ans = [][]string{}

// // 	cols = make(map[int]bool, n)
// // 	pie = make(map[int]bool, n)
// // 	na = make(map[int]bool, n)

// // 	dfs(n, 0, []int{})
// // 	return ans
// // }

// // func dfs(n int, row int, path []int) {
// // 	if row >= n {
// // 		ans = append(ans, generateBoard(path, n))
// // 		return
// // 	}

// // 	for col := 0; col < n; col++ {
// // 		if cols[col] || pie[row + col] || na[row - col] {
// // 			continue
// // 		}
// // 		cols[col] = true
// // 		pie[row + col] = true
// // 		na[row - col] = true
// // 		dfs(n, row + 1, append(path, col))
// // 		cols[col] = false
// // 		pie[row + col] = false
// // 		na[row - col] = false
// // 	}
// // 	return
// // }

//bits
// var solutions [][]string

// func solveNQueens(n int) [][]string {
// 	solutions = [][]string{}
// 	queens := make([]int,n)
// 	for i := 0; i < n; i++ {
// 		queens[i] = -1
// 	}
// 	solve(queens,n,0,0,0,0)
// 	return solutions
// }

// func solve(queens []int,n,row,columns,diagonals1,diagonals2 int) {
// 	if row == n {
// 		board := generateBoard(queens,n)
// 		solutions = append(solutions, board)
// 		return
// 	}

// 	availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
// 	for availablePositions != 0 {
// 		position := availablePositions & - availablePositions
// 		availablePositions = availablePositions & (availablePositions - 1)
// 		column := bits.OnesCount(uint(position - 1))

// 		queens[row] = column
// 		solve(queens, n, row + 1, columns | position, (diagonals1 | position) >> 1, (diagonals2 | position) << 1 )
// 		queens[row] = -1
// 	}
// }

// func generateBoard(queens []int,n int) []string {
// 	res := []string{}

// 	for _, col := range queens {
// 		res = append(res, strings.Repeat(".", col) + "Q" + strings.Repeat(".", n - col - 1))
// 	}

// 	return res
// }

//bit mask
var result [][]string

func solveNQueens(n int) [][]string {
	result = [][]string{}

	board := make([][]byte,n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	helper(board,0,0,0,0)

	return result
}

func helper(board [][]byte,row,colBit,rightBit,leftBit int) {
	n := len(board)

	if row >= n {
		dts := make([]string, n)

		for i,w := range board {
			dts[i] = string(w)
		}

		result = append(result, dts)

		return
	}

	for i := 0; i < n; i++ {
		pick := 1 << (n - i - 1) 

		if pick & colBit == 0 && pick & rightBit == 0 && pick & leftBit == 0 {
			board[row][i] = 'Q'
			helper(board, row + 1, colBit | pick,(rightBit | pick) >> 1,(leftBit | pick) << 1)
			board[row][i] = '.'
		}
	}
}


// @lc code=end

