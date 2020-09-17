/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package main
// @lc code=start
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left) >> 1

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}

	return nums[left]
}
// @lc code=end

