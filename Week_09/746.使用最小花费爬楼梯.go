/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */
 package main
// @lc code=start
func minCostClimbingStairs(cost []int) int {
	//dp[i]	= min(dp[i - 1],dp[i - 2]) + cost[i]
	n := len(cost)

	dp := make([]int,n + 1)


	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i - 1] + cost[i - 1],dp[i - 2] + cost[i - 2])
	}

	return dp[n]
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

