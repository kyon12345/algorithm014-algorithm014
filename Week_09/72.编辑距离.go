/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package main
// @lc code=start
func minDistance(word1 string, word2 string) int {
	//dp[i][j] the edit distance between word[1...i] and word2[1...j]
	//two-end bfs
	m,n := len(word1),len(word2)

	dp := make([][]int,m + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n + 1)
	}

	dp[0][0] = 0
	
	for i := 1;i <= m; i ++ {
		dp[i][0] = i
	}

	for i := 1; i < n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1;j <= n; j ++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}
			dp[i][j] = min(dp[i][j],dp[i][j - 1] + 1)
			dp[i][j] = min(dp[i][j],dp[i - 1][j] + 1)
		}
	}

	return dp[m][n]
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}  
// @lc code=end

