/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package main

// @lc code=start
// func uniquePaths(m int, n int) int {
// 	//dp[m][n]
// 	//dp[0][i] = 1 dp[i][0] = 1
// 	//dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
// 	dp := make([][]int, m)

// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 		dp[i][0] = 1
// 	}

// 	for j := 0; j < n; j++ {
// 		dp[0][j] = 1
// 	}

// 	for i := 1; i < m; i++ {
// 		for j := 1; j < n; j++ {
// 			dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 		}
// 	}

// 	return dp[m-1][n-1]
// }

//滚动数组
func uniquePaths(m int, n int) int {
	//dp[m]
	//对角相加
	//dp[m] += dp[m - 1]
	//base case dp[0] = 1
	dp := make([]int, m)

	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[m-1]
}

// @lc code=end
