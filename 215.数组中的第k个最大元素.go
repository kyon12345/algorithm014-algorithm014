/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

import (
	"math/rand"
	"text/template/parse"
)

// @lc code=start
// func findKthLargest(nums []int, k int) int {
// 	length, left, right := len(nums), 0, len(nums)-1

// 	target := length - k

// 	for {
// 		index := partition(nums, left, right)
// 		if index == target {
// 			return nums[index]
// 			//1 2 3 4 5
			//  i t
			//向右边搜索,因为比这个数大的太多了
// 		} else if index < target {
// 			left = index + 1
// 		} else {
// 			right = index - 1
// 		}
// 	}
// }

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
// func partition(nums []int, left, right int) int {
// 	if right > left {
// 		randomIndex := left + 1 + rand.Intn(right-left)
// 		nums[randomIndex], nums[left] = nums[left], nums[randomIndex]
// 	}

// 	pivot := nums[left]
// 	j := left
// 	for i := left + 1; i <= right; i++ {
// 		if nums[i] < pivot {
// 			j++
// 			nums[j], nums[i] = nums[i], nums[j]
// 		}
// 	}

// 	nums[left], nums[j] = nums[j], nums[left]
// 	return j
// }
// @lc code=end
