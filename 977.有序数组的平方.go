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
	i, j := 0, n-1
	for p := n - 1; p >= 0; p-- {
		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
	}

	return res
}

// @lc code=end
