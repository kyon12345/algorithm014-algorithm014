/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
var res bool

func exist(board [][]byte, word string) bool {
	res = false
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfsWordSearch(board, i, j, "", word)
		}
	}

	return res
}

func dfsWordSearch(board [][]byte, i, j int, s, word string) {
	c := board[i][j]

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '#' {
		return
	}

	if s == word {
		res = true
	}

	board[i][j] = '#'

	dfsWordSearch(board, i+1, j, s+string(c), word)
	dfsWordSearch(board, i-1, j, s+string(c), word)
	dfsWordSearch(board, i, j+1, s+string(c), word)
	dfsWordSearch(board, i, j-1, s+string(c), word)

	board[i][j] = c
}

// @lc code=end
