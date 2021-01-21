/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
package main

// @lc code=start
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < m+1; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = dp[i][j-1]
			if s[j-1] == t[i-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[n][m]
}

// @lc code=end
