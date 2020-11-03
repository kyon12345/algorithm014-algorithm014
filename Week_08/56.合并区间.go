/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package main

import "sort"

// @lc code=start

func merge(intervals [][]int) [][]int {
	res := make([][]int,0)
	
	if len(intervals) < 1 {
		return res
	}

	sort.Slice(intervals, func (i,j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged) - 1]
		c := intervals[i]

		if m[1] < c[0] {
			merged = append(merged, c)
			contin  ue
		}

		if m[1] < c[1] {
			m[1] = c[1]
		}
	}

	return merged
}

// @lc code=end
