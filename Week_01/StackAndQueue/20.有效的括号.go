/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
func isValid(s string) bool {
	//栈,O(n) O(n)
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	stack := make([]rune, 0)

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if len(stack) > 0 && pairs[r] == stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

// @lc code=end
