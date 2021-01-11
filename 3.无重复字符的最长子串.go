/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	//sliding window
	m := [256]int{}

	for i, _ := range m {
		m[i]--
	}

	start, res := -1, 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] > start {
			start = m[s[i]]
		}

		m[s[i]] = i

		if i-start > res {
			res = i - start
		}
	}

	return res
}

// @lc code=end
