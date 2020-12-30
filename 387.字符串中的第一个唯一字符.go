/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
package main

// @lc code=start
func firstUniqChar(s string) int {
	m := [26]int{}

	for i, w := range s {
		m[w-'a'] = i
	}

	for i, w := range s {
		if m[w-'a'] == i {
			return i
		} else {
			m[w-'a'] = -1
		}
	}

	return -1
}

// @lc code=end
