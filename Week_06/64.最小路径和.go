/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
package main

// @lc code=start
func minPathSum(grid [][]int) int {
	//grid[i][j] = min(grid[i -  1][j],grid [i][j - 1]) + grid[i][j]
	//只能从左或者从上来,比较那个值更小就是最佳路径
	//在实际使用中不建议直接修改原数组
	//[[1,2,5],[3,2,1]]
	// if len(grid) == 0 {
	// 	return 0
	// }

	// for i := 0; i < len(grid); i++ {
	// 	for j := 0; j < len(grid[0]); j++ {
	// 		if i == 0 && j == 0 {
	// 			continue
	// 		}

	// 		if i == 0 {
	// 			grid[i][j] += grid[i][j - 1]
	// 		} else if j == 0 {
	// 			grid[i][j] += grid[i - 1][j]
	// 		} else {
	// 			grid[i][j] += min(grid[i - 1][j],grid[i][j - 1])
	// 		}

	// 	}
	// }

	// return grid[len(grid) - 1][len(grid[0]) - 1]

	if len(grid) < 1 {
		return 0
	}

	row, col := len(grid), len(grid[0])

	dp := make([]int, col+1)

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			if i == 1 {
				dp[j] = dp[j-1] + grid[i-1][j-1]
			} else if j == 1 {
				dp[j] += grid[i-1][j-1]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i-1][j-1]
			}
		}
	}

	return dp[col]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
