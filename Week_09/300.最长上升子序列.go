/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */
package main

// @lc code=start
func lengthOfLIS(nums []int) int {
	//dp[i]	= max(dp[j]) + 1  0 <= j < i && nums[i] > nums[j]

	//special
	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums))

	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(dp[i],res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
