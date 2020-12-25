/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
package main

import "math"

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)

	res := make([]int, n)

	l,r := 0,n - 1
	for p := n - 1; p >= 0; p-- {
		if abs(nums[l]) < abs(nums[r]) {
			res[p] = nums[r] * nums[r]
			r --
		} else {
			res[p] = nums[l] * nums[l]
			l ++
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
