/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */
package main

import "sort"

// @lc code=start
func findContentChildren(g []int, s []int) int {
	//优先使用小的饼干来满足孩子,使用双指针来优化时间复杂度
	sort.Ints(g)
	sort.Ints(s)

	i,j := 0,0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i ++
		}
		j ++
	}

	return i
}
// @lc code=end

