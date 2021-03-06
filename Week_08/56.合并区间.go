/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package main

import "sort"

// @lc code=start

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := range intervals {
		m := merged[len(merged) - 1]
		c := intervals[i]

		if m[1] < c[0] {
			merged = append(merged, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}

	return merged
}

// @lc code=end
