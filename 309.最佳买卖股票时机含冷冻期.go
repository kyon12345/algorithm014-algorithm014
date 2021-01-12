/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package main

import "math"

// @lc code=start
// func maxProfit(prices []int) int {
// 	//T[i][k][0] = max(T[i - 1][k][0],T[i - 1][k][1] + prices[i])
// 	//如果要买入股票,不能在前一天买入
// 	//T[i][k][1] = max(T[i - 1][k][1].T[i - 2][k][0] - prices[i])
// 	if len(prices) < 2 {
// 		return 0
// 	}

// 	n := len(prices)

// 	dp := make([][2]int, n)

// 	dp[0][1] = -prices[0]
// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// 		if i >= 2 {
// 			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
// 		} else {
// 			dp[i][1] = max(dp[i-1][1], -prices[i])
// 		}
// 	}

// 	return dp[n-1][0]
// }

func maxProfit(prices []int) int {
	sell,pre_sell,buy,pre_buy := 0,0,math.MinInt32,0

	for _, p := range prices {
		pre_buy = buy
		buy = max(pre_sell - p,pre_buy)
		pre_sell = sell
		sell = max(pre_sell,pre_buy + p)
	}

	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
