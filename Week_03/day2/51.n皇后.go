/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */
 package queen

// @lc code=start
func solveNQueens(n int) [][]string {
	solution := make([][]string, 0)
	cUsed, dUsed := make([]bool, n), make([]bool, (n*2-1)*2)
	backtrack(&solution, buildNxNBoard(n), cUsed, dUsed, n, 0)
	return solution
}

func backtrack(solution *[][]string, board [][]byte, cUsed, dUsed []bool, n, y int) {
	if y == n {
		boardCopy := make([]string, n)
		for i := range boardCopy {
			boardCopy[i] = string(board[i])
		}
		*solution = append(*solution, boardCopy)
	} else {
		for x := 0; x < n; x++ {
			if !cUsed[x] && !dUsed[y-x+n-1] && !dUsed[y+x+n*2-1] {
				cUsed[x], dUsed[y-x+n-1], dUsed[y+x+n*2-1] = true, true, true
				board[y][x] = 'Q'
				backtrack(solution, board, cUsed, dUsed, n, y+1)
				board[y][x] = '.'
				cUsed[x], dUsed[y-x+n-1], dUsed[y+x+n*2-1] = false, false, false
			}
		}
	}
}

func buildNxNBoard(n int) [][]byte {
	result := make([][]byte, n)
	str := make([]byte, n)
	for i := range str {
		str[i] = '.'
	}
	for i := range result {
		result[i] = append([]byte{}, str...)
	}
	return result
}
// @lc code=end

