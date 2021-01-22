/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

import (
	"math/rand"
)

// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)

	target := n - k

	lo, hi := 0, n-1
	for {
		index := partition(nums, lo, hi)

		if index == target {
			return nums[index]
		} else if index < target {
			lo = index + 1
		} else {
			hi = index - 1
		}
	}
}

//因为选取的是固定的指针所以对一些特殊的测试用例会导致复杂度上升为N^2
// func partition(nums []int, left, right int) int {
// 	pivot := nums[left]
// 	j := left

// 	for i := left + 1; i <= right; i++ {
// 		if nums[i] < pivot {
// 			j++
// 			nums[i], nums[j] = nums[j], nums[i]
// 		}
// 	}

// 	nums[j], nums[left] = nums[left], nums[j]

// 	return j
// }

//随机取一个哨兵节点
//O(N)
func partition(nums []int, left, right int) int {
	if left < right {
		ranIndex := left + rand.Intn(right-left) + 1
		nums[ranIndex], nums[left] = nums[left], nums[ranIndex]
	}

	pivot := nums[left]
	j := left

	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[j], nums[left] = nums[left], nums[j]

	return j
}

// @lc code=end
