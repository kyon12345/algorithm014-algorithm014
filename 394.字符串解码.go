/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

import "strings"

// @lc code=start
//stack
func decodeString(s string) string {
	stackNums := make([]int,0)
	stackStr := make([]string,0)

	var res string
	var num int

	for i := 0; i < len(s); i++ {
		switch  {
		case s[i] <= '9' && s[i] >= '0':
			num = num * 10 + int(s[i]) - '0' 
		case s[i] == '[':
			stackNums = append(stackNums, num)
			stackStr = append(stackStr, res)
			num = 0
			res = ""
		case s[i] == ']':
			tmp := stackStr[len(stackStr) - 1]
			stackStr = stackStr[:len(stackStr) - 1]
			count := stackNums[len(stackNums) - 1]
			stackNums = stackNums[:len(stackNums) - 1]
			res = tmp + strings.Repeat(res, count)
		default:
			res += string(s[i])
		}
	}

	return res
}
// @lc code=end

