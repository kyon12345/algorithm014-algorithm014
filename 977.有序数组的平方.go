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
	
	res := make([]int,n)

	lo,hi := 0,n - 1
	for p := n - 1; p >= 0; p-- {
		if math.Abs(float64(nums[lo])) < math.Abs(float64(nums[hi])) {
			res[p] = nums[hi] * nums[hi]
			hi --
		} else {
			res[p] = nums[lo] * nums[lo]
			lo ++
		}
	}

	return res
}

// @lc code=end
