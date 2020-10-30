/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package main

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([]int, m)

	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m-1]
}

// @lc code=end
