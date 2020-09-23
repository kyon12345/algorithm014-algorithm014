/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

import "math"

// @lc code=start
func coinChange(coins []int, amount int) int {
	// state: amount
	// dp(amount): fewest number of coins to make amount
	// choice: use each of the coins
	// base: dp(0) = 0

	// dp(amount) = min(dp(amount-coin), for each coin)+1

	dp := make([]int, amount+1, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = math.MaxInt32
        for _, c := range coins {
            if i - c >= 0 {
                dp[i] = min(dp[i-c]+1, dp[i])
            }
        }
    }
    if dp[amount] == math.MaxInt32 {
        return -1
    }
    
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
// @lc code=end

