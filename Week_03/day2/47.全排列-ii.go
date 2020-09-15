/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package main

import "sort"

// @lc code=start
//O(N * N!) O(N * N!)
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	res := make([][]int, 0)

	backtrackPermute2(nums, []int{}, &res)

	return res
}

func backtrackPermute2(nums, prev []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		backtrackPermute2(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(prev, nums[i]), res)
	}
}

// @lc code=end
