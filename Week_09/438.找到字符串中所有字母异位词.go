/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package main
// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	have,need := [26]int{},[26]int{}

	for i := range p {
		have[s[i] - 'a'] ++
		need[p[i] - 'a'] ++
	}

	res := make([]int,0)
	head := 0
	for tail := len(p); tail < len(s); tail++ {
		if have == need {
			res = append(res, head)
		}

		have[s[head] - 'a'] --
		have[s[tail] - 'a'] ++

		head ++
	}

	if have == need {
		res = append(res, head)
	}

	return res
}
// @lc code=end

