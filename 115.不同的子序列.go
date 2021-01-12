/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
package main

// @lc code=start
func numDistinct(s string, t string) int {
	n, m := len(s), len(t)

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}

	for j := 1; j < m+1; j++ {
		dp[0][j] = 0
	}

	for j := 1; j < m+1; j++ {
		for i := 1; i < n+1; i++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[n][m]
}

// @lc code=end
