/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
// package main
package main
// @lc code=start
func longestPalindrome(s string) string {
	//dp[i][j] = dp[i + 1][j - 1] 
	if len(s) < 2 {
		return s
	}

	n := len(s)
	dp := make([][]bool,n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	begin,maxL := 0,1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else if i - j < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i - 1][j + 1]
			}

			if dp[i][j] && maxL < i - j + 1 {
				begin = j
				maxL = i - j + 1
			}
		}
	}

	return s[begin:begin + maxL]
}

// @lc code=end
