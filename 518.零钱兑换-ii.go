/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */
package main

// @lc code=start
func change(amount int, coins []int) int {
	//不用第i枚硬币只使用i - 1枚硬币的解法
	//dp[i][j] = dp[i - 1][j] + dp[i][j - coins[i - 1]]
	dp := make([]int,amount + 1)

	dp[0] = 1
	for _, c := range coins {
		for i := c; i < amount + 1; i++ {
			dp[i] += dp[i - c]
		}
	}

	return dp[amount]
}	

// @lc code=end
