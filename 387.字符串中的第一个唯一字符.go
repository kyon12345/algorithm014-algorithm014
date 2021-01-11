/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
package main

// @lc code=start
func firstUniqChar(s string) int {
	m := [26]int{}

	for i, c := range s {
		m[c-'a'] = i
	}

	for i, c := range s {
		if i == m[c-'a'] {
			return i
		} else {
			m[c-'a'] = -1
		}
	}

	return -1
}

// @lc code=end
