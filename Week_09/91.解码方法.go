/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package main
// @lc code=start
func numDecodings(s string) int {
	dp := make([]int,len(s) + 1)

	dp[0] = 1
	dp[1] = 0

	if s[0] != '0' {
		dp[1] = 1
	}
	
	for	i := 2;i <= len(s);i ++ {
		oneDig := s[i - 1] - '0'
		twoDig := (s[i - 2] - '0') * 10 + oneDig

		if oneDig >= 1 {
			dp[i] += dp[i - 1]
		}

		if twoDig >= 10 && twoDig <= 26 {
			dp[i] += dp[i - 2]
		}
	}

	return dp[len(s)]
}
// @lc code=end

