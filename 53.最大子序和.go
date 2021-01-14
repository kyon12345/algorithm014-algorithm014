/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+nums[i-1] {
			nums[i] = nums[i] + nums[i-1]
		}

		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}

// @lc code=end
