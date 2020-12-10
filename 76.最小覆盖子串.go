/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

import "math"

// @lc code=start
func minWindow(s string, t string) string {
	m := map[byte]int{}

	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	start, end := 0, 0
	counter := len(t)
	minStart, minlen := 0, math.MaxInt32
	size := len(s)

	for end < size {
		if m[s[end]] > 0 {
			counter--
		}

		//Decrease m[s[end]],if does not exist in t,m [s[end]] will be negative
		m[s[end]]--
		end++

		//when we find valid window,we start to find smaller window
		for counter == 0 {
			if end-start < minlen {
				minStart = start
				minlen = end - start
			}

			m[s[start]]++
			//when char exists in t,increase counter
			if m[s[start]] > 0 {
				counter++
			}
			start++
		}
	}

	if minlen != math.MaxInt32 {
		return s[minStart : minStart+minlen]
	}

	return ""
}

// @lc code=end
