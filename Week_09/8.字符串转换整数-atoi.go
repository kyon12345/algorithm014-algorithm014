/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
package main

import "math"

// @lc code=start
func myAtoi(str string) int {
	res,sign,len,idx := 0,1,len(str),0

	//skip leading sapces
	for idx < len && (str[idx] == ' ' || str[idx] == '\t') {
		idx ++
	}

	// +/- sign
	if idx < len {
		if str[idx] == '+' {
			sign = 1
			idx ++
		} else if str[idx] == '-' {
			sign = -1
			idx ++
		}
	}

	//Numbers
	for idx < len && str[idx] >= '0' && str[idx] <= '9' {
		res = res * 10 + int(str[idx]) - '0'
		if sign * res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign * res < math.MinInt32 {
			return math.MinInt32
		}
		idx ++
	}

	return res * sign
}
// @lc code=end

