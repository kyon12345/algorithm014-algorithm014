/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */
package main
// @lc code=start
// func longestCommonSubsequence(text1 string, text2 string) int {
// 	//If S1[-1] != S2[-1]: LCS[s1, s2] = Max(LCS[s1-1, s2], LCS[s1, s2-1]) 
// 	// If S1[-1] == S2[-1]: LCS[s1, s2] = LCS[s1-1, s2-1] + 1
// 	dp := make([][]int,len(text1) + 1)

// 	for i := 0; i <= len(text1); i++ {
// 		dp[i] = make([]int, len(text2) + 1)
// 	}

// 	for i := 1; i <= len(text1); i++ {
// 		for j := 1; j <= len(text2); j++ {
// 			if text1[i - 1] == text2[j - 1] {
// 				dp[i][j] = dp[i - 1][j - 1] + 1
// 			} else {
// 				dp[i][j] = max(dp[i - 1][j],dp[i][j - 1])
// 			}
// 		}
// 	}
	
// 	return dp[len(text1)][len(text2)]
// }

//降维
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2) + 1)

	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j]

			if text1[i - 1] == text2[j - 1] {
				dp[j] = last + 1
			} else {
				dp[j] = max(tmp,dp[j - 1])
			}

			last = tmp
		}
	}

	return dp[len(text2)]	
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

