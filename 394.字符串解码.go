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
	numStack := []int{}
	strStack := []string{}

	res := ""
	num := 0

	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			num += num*10 + int(c) - '0'
		case c == '[':
			numStack = append(numStack, num)
			strStack = append(strStack, res)
			res = ""
			num = 0
		case c == ']':
			tmpNum := numStack[len(numStack)-1]
			tmpS := strStack[len(strStack)-1]
			numStack = numStack[:len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
			res = tmpS + strings.Repeat(res, tmpNum)
		default:
			res += string(c)
		}
	}

	return res
}

// @lc code=end
