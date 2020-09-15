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

		res := make([][]int,0)

		sort.Ints(nums)

		backtrackPermuteUnique(nums, []int{}, &res)

		return res
	}

	func backtrackPermuteUnique(nums,prev []int,res *[][]int) {
		if len(nums) == 0 {
			// *res = append(*res,append([]int{},prev...))
			//这样就无法通过
			*res = append(*res,prev)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i >0 && nums[i] == nums[i - 1] {
				continue
			}

			backtrackPermuteUnique(append(append([]int{},nums[:i]...),nums[i + 1:]...),
			append(prev, nums[i]), res)
		}
	}
// @lc code=end

