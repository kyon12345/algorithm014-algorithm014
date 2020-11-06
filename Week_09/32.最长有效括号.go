/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package main

import "internal/poll"

// @lc code=start

//stack O(n) O(n)
func longestValidParentheses(s string) int {
	//[-1,idx...] max = max(idx - newTop)
	stack := []int{-1}

	maxVal := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if i - stack[len(stack) - 1] > maxVal {
					maxVal = i - stack[len(stack) - 1]
				}
			}
		}
	}

	return maxVal
}

//dp
// func longestValidParentheses(s string) int {
// 	dp := make([]int,len(s))

// 	maxVal := 0

// 	for i := 1; i < len(s); i++ {
// 		if s[i] == ')' {
// 			if s[i - 1] == '(' {
// 				dp[i] = 2
// 				if i - 2 >= 0 {
// 					dp[i] += dp[i - 2]
// 				}
// 			} else if dp[i - 1] > 0 {
// 				if i - dp[i - 1] - 1 >= 0 && s[i - dp[i - 1] - 1] == '(' {
// 					dp[i] = dp[i - 1] + 2
// 					if i - dp[i - 1] - 2 >= 0 {
// 						dp[i] += dp[i - dp[i - 1] - 2]
// 					}
// 				}
// 			}
// 		}
		
// 		if dp[i] > maxVal {
// 			maxVal = dp[i]
// 		}
// 	}

// 	return maxVal
// }

// @lc code=end
