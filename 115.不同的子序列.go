/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
package main

// @lc code=start
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			dp[i][j] = dp[i][j-1]
			if t[i-1] == s[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
