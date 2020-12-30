/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */
package main

import "strconv"

// @lc code=start
func diffWaysToCompute(input string) []int {
	if isDigit(input) {
		res, _ := strconv.Atoi(input)

		return []int{res}
	}

	res := []int{}

	for index, w := range input {
		if w == '+' || w == '-' || w == '*' {
			left, right := diffWaysToCompute(input[:index]), diffWaysToCompute(input[index+1:])

			var addNum int
			for _, l := range left {
				for _, r := range right {
					if w == '+' {
						addNum = l + r
					} else if w == '-' {
						addNum = l - r
					} else {
						addNum = l * r
					}

					res = append(res, addNum)
				}
			}
		}
	}

	return res
}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)

	if err != nil {
		return false
	}

	return true
}

// @lc code=end
