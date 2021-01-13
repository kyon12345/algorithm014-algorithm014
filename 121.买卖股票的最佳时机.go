/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
// func maxProfit(prices []int) int {
// 	if len(prices) < 2 {
// 		return 0
// 	}

// 	n := len(prices)

// 	dp := make([][2]int, n)

// 	dp[0][0] = 0
// 	dp[0][1] = -prices[0]

// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// 		dp[i][1] = max(dp[i-1][1], -prices[i])
// 	}

// 	return dp[n-1][0]
// }

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	hold,cash := -prices[0],0
	for i := 1; i < len(prices); i++ {
		hold = max(hold,-prices[i])
		cash = max(cash, hold + prices[i])
	}

	return cash
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
