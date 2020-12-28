/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

import "fmt"

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return kSelect(nums, 0, len(nums)-1, k)
}

func kSelect(nums []int, left int, right int, k int) int {
	if left == right {
		return nums[left]
	}

	p := right
	l := left
	r := right - 1
	for l <= r {
		for l <= r && nums[l] >= nums[p] {
			l++
		}

		for l <= r && nums[r] < nums[p] {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[l], nums[p] = nums[p], nums[l]
	p = l

	if p > k-1 {
		return kSelect(nums, left, p-1, k)
	} else if p < k-1 {
		return kSelect(nums, p+1, right, k)
	}

	fmt.Println(nums)

	return nums[p]
}

// @lc code=end
