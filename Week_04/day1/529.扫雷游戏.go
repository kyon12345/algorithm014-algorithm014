/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */
package main
// @lc code=start

var dirX = []int{1,-1,1,-1,-1,1,0,0}
var dirY = []int{0,0,1,-1,1,-1,1,-1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfsBoard(board,x,y)
	}

	return board
}

func dfsBoard(board [][]byte,x,y int) {
	cnt := 0 

	for i := 0; i < 8; i++ {
		tx,ty := x + dirX[i],y + dirY[i]

		if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) {
			continue
		}

		if board[tx][ty] == 'M' {
			cnt ++
		}
	}

	if cnt > 0 {
		board[x][y] = byte('0' + cnt)
	} else {
		board[x][y] = 'B'

		for i := 0; i < 8; i++ {
			tx,ty := x + dirX[i],y + dirY[i]

			if tx < 0 || tx >= len(board) || ty <0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}

			dfsBoard(board, tx, ty)
		}
	}
}

// @lc code=end

