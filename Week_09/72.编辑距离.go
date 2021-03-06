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
	m, n := len(word1), len(word2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < m+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if word1[j-1] == word2[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			dp[i][j] = min(dp[i][j], min(dp[i-1][j], dp[i][j-1])+1)
		}
	}

	return dp[n][m]
}

//dfs 时间复杂度爆炸
// func minDistance(word1 string, word2 string) int  {
// 	if len(word1) == 0 && len(word2) == 0 {
// 		return 0
// 	}

// 	if len(word1) == 0 {
// 		return len(word2)
// 	}

// 	if len(word2) == 0 {
// 		return len(word1)
// 	}

// 	if word1[0] == word2[0] {
// 		return minDistance(word1[1:], word2[1:])
// 	}

// 	insert := 1 + minDistance(word1, word2[1:])
// 	delete := 1 + minDistance(word1[1:], word2)
// 	replace := 1 + minDistance(word1[1:], word2[1:])

// 	return min(insert,min(delete,replace))
// }

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
