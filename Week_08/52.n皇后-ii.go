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

	helper(n,0,0,0,0)

	return count
}

func helper(n, row, col, left, right int) {
	if row == n {
		count++
		return
	}

	// ^(0001 | 0010 | 0000) & (1111) = 1100
	availablePositions := ^(col | left | right) & (1<<n - 1)
	for availablePositions > 0 {
		pick := availablePositions & -availablePositions
		helper(n, row+1, col|pick, (left|pick)<<1, (right|pick)>>1)
		availablePositions &= availablePositions - 1
	}
}

// @lc code=end
