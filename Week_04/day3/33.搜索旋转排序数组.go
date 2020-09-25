/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package main

import "math"

// @lc code=start
func search(nums []int, target int) int {
	//二分查找
	//首先思考没有中断的数组是怎么寻找的
	

	l,r := 0, len(nums) - 1

	for l <= r {
		mid := l + (r - l) / 2

		if nums[mid] == target {
			return mid
		}
	
		//如果mid > l 说明在左半边
		if nums[mid] >= nums[l] {
			//继续在左半边寻找(单调的)
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//如果在右半边,且介于mid和右边界之间,那我们的左边界应该右移来取到目标值(偏右)
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
// @lc code=end

