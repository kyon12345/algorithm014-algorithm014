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
		m[c] ++
	}

	minStart,minlen := 0,math.MaxInt32
	end,counter,start := 0,len(t),0

	for end < len(s) {
		c := s[end]

		if m[c] > 0 {
			counter --
		}

		m[c] --
		end ++

		for counter == 0 {
			if end - start < minlen {
				minStart = start
				minlen = end - start
			}

			m[s[start]] ++

			if m[s[start]] > 0 {
				counter ++
			}

			start ++
		}
	}

	if minlen != math.MaxInt32 {
		return s[minStart:minStart + minlen]
	} else {
		return ""
	}
}

// @lc code=end
