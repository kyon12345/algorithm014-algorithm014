/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */
package main
// @lc code=start
var count int
func totalNQueens(n int) int {
    //special
    if n <= 0 {
        return count
    }
    //init
    count = 0 
    //begin dfs
    dfs(n,0,0,0,0)

    return count
}

func dfs(n,row,colBit,leftBit,rightBit int) {
    //termination condition
    if row == n {
        count ++
        return
    }

    //process
    avaliablePosition := ^(colBit | leftBit | rightBit) & ((1 << n) - 1)

    //loop until got no position
    for avaliablePosition > 0 {
        //got the last 1
        pick := avaliablePosition & -avaliablePosition
        //drill down
        dfs(n,row + 1,colBit | pick,(leftBit | pick) << 1,(rightBit | pick) >> 1)
        //remove last 1
        avaliablePosition &= avaliablePosition - 1
    }
}
// @lc code=end

