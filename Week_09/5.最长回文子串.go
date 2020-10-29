/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

import "golang.org/x/net/ipv6"

// @lc code=start
func longestPalindrome(s string) string {
	l := len(s)

	if l < 2 {
		return s
	}

	dp := make([][]bool, l)

	//init chart
	for i := range dp {
		dp[i] = make([]bool, l)
	}

	maxLength,begin := 1,0
	
	chars := []byte(s)

	for j := 1;j < l;j ++ {
		for i := 0;i < j;i ++ {
			if chars[i] != chars[j] {
				dp[i][j] = false
			} else if j - i < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i + 1][j - 1] 
			}

			if dp[i][j] && j - i + 1 > maxLength {
				maxLength = j - i + 1
				begin = i
			}
		}
	}

	return string(chars[begin:begin + maxLength])
}

// @lc code=end
