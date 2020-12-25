/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && isFound(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func isFound(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}

	if word[index] != board[i][j] {
		return false
	}

	c := board[i][j]
	board[i][j] = '#'
	if isFound(board, i+1, j, word, index+1) ||
		isFound(board, i-1, j, word, index+1) ||
		isFound(board, i, j-1, word, index+1) ||
		isFound(board, i, j+1, word, index+1) {
		return true
	}

	board[i][j] = c
	return false
}

// @lc code=end
