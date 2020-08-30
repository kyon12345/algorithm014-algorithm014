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

	ans := make([][]int,0)

	sort.Ints(nums)

	backtrack2(nums,nil,&ans)

	return ans
}

func backtrack2(nums,prev []int,ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans,append([]int{},prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		backtrack2(append(append([]int{},nums[:i]...),nums[i + 1:]...),
		append(prev,nums[i]),ans)
	}
}
// @lc code=end

