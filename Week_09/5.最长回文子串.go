/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
// package main


// @lc code=start
func longestPalindrome(s string) string {
	l := len(s)
	
	if l < 2 {
		return s
	}

	dp := make([][]bool,l)

	for i := range dp {
		dp[i] = make([]bool, l)
	}

	idx,maxLength := 0,1

	for j := 0; j < l; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else if j - i < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i + 1][j - 1]
			}

			if dp[i][j] && j - i + 1 > maxLength {
				idx = i
				maxLength = j - i + 1
			}
		}
	}

	return s[idx:idx + maxLength]
}

// @lc code=end
