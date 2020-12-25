/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */
package main

// @lc code=start
func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo)>>1

		if nums[mid] == nums[hi] {
			hi--
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return nums[lo]
}

// @lc code=end
