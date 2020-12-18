/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */
package main
// @lc code=start
var count int
func totalNQueens(n int) int {
    count = 0

    dfs(n,0,0,0,0)

    return count
}

func dfs(n,row,col,right,left int) {
    if row == n {
        count ++
        return
    }

    availablePositions := ^(col | right | left) & (1 << n - 1)

    for availablePositions > 0 {
        pick := availablePositions & -availablePositions
        dfs(n, row + 1, col | pick, (right | pick) >> 1, (left | pick) << 1)
        availablePositions &= availablePositions - 1
    }
}
// @lc code=end
