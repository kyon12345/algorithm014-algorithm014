/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package main

// @lc code=start
//O(log n) O(1)
func findMin(nums []int) int {
	//本质是寻找一个变化点,这个变化点满足
	//1.变化点左边的所有值都大于数组的第一个元素
	//2.变化点右边的所有值都小于数组的第一个元素
	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return nums[0]
	}

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		//如果最小值在右边,就在右半边进行搜索
		if nums[0] < nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

// @lc code=end
