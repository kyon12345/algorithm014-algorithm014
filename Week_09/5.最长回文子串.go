/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
// package main
package main
// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool,len(s))

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	idx,maxLength := 0,1
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] != s[i] {
				dp[i][j] = false
			} else if i - j < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i - 1][j + 1]
			}

			if dp[i][j]	&& i - j + 1 > maxLength {
				idx = j
				maxLength = i - j + 1
			}			
		}
	}

	return s[idx:idx + maxLength]
}

// @lc code=end
