/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */
package main

import "sort"

// @lc code=start
func findContentChildren(g []int, s []int) int {
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

