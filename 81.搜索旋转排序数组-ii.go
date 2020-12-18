/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */
package main

// @lc code=start
//O(logn) worst O(n)
func search(nums []int, target int) bool {
	lo,hi := 0,len(nums) - 1
	

	//(mid,hi],[lo,mid)
	for lo <= hi {
		mid := lo + (hi - lo) >> 1

		if target == nums[mid] {
			return true
		}

		if nums[lo] == nums[mid] && nums[hi] == nums[mid] {
			lo ++ 
			hi --
		} else if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return false

}

// @lc code=end
