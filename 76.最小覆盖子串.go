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

	counter,end := len(t),0
	start,minLen := 0,math.MaxInt32
	minStart := 0
	
	for end < len(s) {
		if m[s[end]] > 0 {
			counter --
		}

		m[s[end]] --
		end ++

		for counter == 0 {
			if minLen > end - start {
				minLen = end - start
				minStart = start
			}

			m[s[start]] ++

			if m[s[start]] > 0 {
				counter ++
			}

			start ++
		}
	}

	if minLen != math.MaxInt32 {
		return s[minStart:minStart + minLen]
	}

	return ""
}

// @lc code=end
