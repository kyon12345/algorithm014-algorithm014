/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package main
// @lc code=start
func findAnagrams(s string, p string) []int {
	res := []int{}

	if len(s) < len(p) {
		return res
	}

	have,need := [26]int{},[26]int{}

	for i := 0; i < len(p); i++ {
		have[s[i] - 'a'] ++
		need[p[i] - 'a'] ++
	}

	head := 0
	for i := len(p); i < len(s); i++ {
		if have == need {
			res = append(res, head)
		}

		have[s[head] - 'a'] --
		have[s[i] - 'a'] ++

		head ++
	}

	if have == need {
		res = append(res, head)
	}

	return res
}
// @lc code=end

