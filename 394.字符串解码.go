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
	strStack := make([]string,0)
	numStack := make([]int, 0)

	var num int
	var res string

	for _, v := range s {
		switch {
		case v <= '9' && v >= '0':
			num = num * 10 + int(v) - '0'
		case v == '[':
			numStack = append(numStack, num)
			strStack = append(strStack, res)
			num = 0
			res = ""
		case v == ']':
			tmp := strStack[len(strStack) - 1]
			strStack = strStack[:len(strStack) - 1]
			count := numStack[len(numStack) - 1]
			numStack = numStack[:len(numStack) - 1]
			res = tmp + strings.Repeat(res, count) 
		default:
			res += string(v)
		}
	}

	return res
}
// @lc code=end

