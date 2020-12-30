/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

import "math"

// @lc code=start
func minWindow(s string, t string) string {
	m := [128]int{}

	for _, c := range t {
		m[c]++
	}

	start, end := 0, 0
	minlen, minstart := math.MaxInt32, 0
	counter := len(t)

	for end < len(s) {
		c := s[end]

		if m[c] > 0 {
			counter--
		}

		m[c]--

		end++

		for counter == 0 {
			if end-start < minlen {
				minlen = end - start
				minstart = start
			}

			m[s[start]]++
			if m[s[start]] > 0 {
				counter++
			}

			start++
		}
	}

	if minlen != math.MaxInt32 {
		return s[minstart : minstart+minlen]
	} else {
		return ""
	}
}

// @lc code=end
