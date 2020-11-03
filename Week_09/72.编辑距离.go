/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package main
// @lc code=start
// func minDistance(word1 string, word2 string) int {
// 	//dp[i][j] the edit distance between word[1...i] and word2[1...j]
// 	//two-end bfs
// 	m,n := len(word1),len(word2)

// 	dp := make([][]int,m + 1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, n + 1)
// 	}

// 	for i := 0; i < m + 1; i++ {
// 		dp[i][0] = i
// 	}

// 	for j := 0; j < n + 1; j++ {
// 		dp[0][j] = j
// 	}

// 	for i := 1; i < m + 1; i++ {
// 		for j := 1; j < n + 1; j++ {
// 			if word1[i - 1] == word2[j - 1] {
// 				dp[i][j] = dp[i - 1][j - 1]
// 			} else {
// 				dp[i][j] = dp[i - 1][j - 1] + 1
// 			}

// 			dp[i][j] = min(dp[i - 1][j] + 1,min(dp[i][j],dp[i][j - 1] + 1))
// 		}
// 	}

// 	return dp[m][n]
// }

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

// func min(a,b int) int {
// 	if a < b {
// 		return a
// 	}

// 	return b
// }  
// @lc code=end

