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

	result := make([]int,0)
	need,have := [26]int{},[26]int{}

	for i := range p {
		need[p[i] - 'a'] ++
		have[s[i] - 'a'] ++
	}

	var tail int
	for i := len(p);i < len(s);i ++ {
		if have == need {
			result = append(result, tail)
		}

		//remove the head and add the tail
		have[s[tail] - 'a'] --
		have[s[i] - 'a'] ++

		tail ++
	}

	if have == need {
		result = append(result, tail)
	}

	return result
}
// @lc code=end

