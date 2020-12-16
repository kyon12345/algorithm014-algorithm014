/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */
package main

// @lc code=start
//O(logn) worst O(n)
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	mid := 0

	for left <= right {	
		mid = (left + right) >> 1
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

// @lc code=end
