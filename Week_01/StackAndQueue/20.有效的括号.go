/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
func isValid(s string) bool {
	//栈,O(n) O(n)
	pairs := map[rune]rune{'}':'{',']':'[',')':'('}

	stack := []rune{}

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && pairs[c] == stack[len(stack) - 1] {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end
