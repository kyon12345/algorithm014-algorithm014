/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */
package main


// @lc code=start

var dirX = []int{1, -1, 1, -1, -1, 1, 0, 0}
var dirY = []int{0, 0, 1, -1, 1, -1, 1, -1}

//O(MN) O(MN)
func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	dfsBoard(board, x, y)

	return board
}

func dfsBoard(board [][]byte, x, y int) {

	cnt := 0

	for i := 0; i < 8; i++ {
		dx, dy := x+dirX[i], y+dirY[i]
		if dx < 0 || dx >= len(board) || dy < 0 || dy >= len(board[0]) {
			continue
		}

		if board[dx][dy] == 'M' {
			cnt++
		}
	}

	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
	} else {
		board[x][y] = 'B'

		for i := 0; i < 8; i++ {
			dx, dy := x+dirX[i], y+dirY[i]

			if dx < 0 || dx >= len(board) || dy < 0 || dy >= len(board[0]) || board[dx][dy] != 'E' {
				continue
			}

			dfsBoard(board, dx, dy)
		}
	}
}

// @lc code=end
