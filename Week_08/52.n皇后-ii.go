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

func dfs(n,row,col,left,right int) {
    if row == n {
        count ++
        return
    }

    avaliablePos := ^(col | left | right) & ((1 << n) - 1)

    for avaliablePos > 0 {
        pick := avaliablePos & -avaliablePos

        dfs(n,row + 1,col | pick,(left | pick) << 1,(right | pick) >> 1)

        avaliablePos &= avaliablePos - 1
    }
}
// @lc code=end

