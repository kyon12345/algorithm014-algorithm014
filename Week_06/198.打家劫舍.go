/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package main
// @lc code=start
// func rob(nums []int) int {
// 	//dp[i]=max(dp[i−2]+nums[i],dp[i−1])
// 	//dp[0] = nums[0]
// 	//d[[1] = max(nums[0],nums[1])
// 	if len(nums) == 0 {
//         return 0
//     }
//     if len(nums) == 1 {
//         return nums[0]
//     }
//     dp := make([]int, len(nums))
//     dp[0] = nums[0]
//     dp[1] = max(nums[0], nums[1])
//     for i := 2; i < len(nums); i++ {
//         dp[i] = max(dp[i-2] + nums[i], dp[i-1])
//     }
//     return dp[len(nums)-1]
// }
func rob(nums []int) int {
	//滚动数组,dp[i] = max(dp[i - 2] + num, dp[i = 1])
	if len(nums) == 0 {
		return 0
	}


	p,q,tmp := 0,0,0

	for _, num := range nums {
		tmp = p
		p = max(p, q + num)
		q = tmp
	}

	return p
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

