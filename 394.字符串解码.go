/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

import "strings"

// @lc code=start
//stack
//2[a2[bc]]
func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0)

	num, res := 0, ""

	for _, c := range s {
		switch {
		case c <= '9' && c >= '0':
			num = num*10 + int(c) - '0'
		case c == '[':
			numStack = append(numStack, num)
			strStack = append(strStack, res)
			num = 0
			res = ""
		case c == ']':
			tmpNum := numStack[len(numStack) - 1]
			tmpStr := strStack[len(strStack) - 1]
			numStack = numStack[:len(numStack) - 1]
			strStack = strStack[:len(strStack) - 1]
			res = tmpStr + strings.Repeat(res, tmpNum)
		default:
			res += string(c)
		}
	}

	return res
}

// @lc code=end
