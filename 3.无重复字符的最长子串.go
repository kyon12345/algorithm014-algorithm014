/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	//sliding window
	dic := [256]int{}

	for i := 0; i < len(dic); i++ {
		dic[i] = -1
	}

	start,maxlen := -1,0

	for i := 0; i < len(s); i++ {
		if dic[s[i]] > start {
			start = dic[s[i]]
		}

		dic[s[i]] = i

		if i - start > maxlen {
			maxlen = i - start
		}
	}

	return maxlen
}
// @lc code=end

