/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package main
// @lc code=start
func numDecodings(s string) int {
	// dp[i] = dp[i - 1] + dp[i - 2]
	//每次只考虑当前字符的前两个字符可能性
	//1.都满足 +2 + 1
	//2.有一个满足 ++ 1
	//都不满足 0,遍历下一个位置
	dp := make([]int,len(s) + 1)

	dp[0] = 1
	dp[1] = 0

	if s[0] != '0' {dp[1] = 1}
	for i := 2; i <= len(s); i++ {
		oneDig := s[i - 1] - '0'
		twoDig := (s[i - 2] - '0')*10 + oneDig

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

