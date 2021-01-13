/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package main

// @lc code=start
//贪心
// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}

// 	res := 0

// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] > prices[i-1] {
// 			res += prices[i] - prices[i-1]
// 		}
// 	}

// 	return res
// }
// //动态规划
// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	if n < 2 {
// 		return 0
// 	}

// 	dp := make([][2]int, n)

// 	dp[0][0] = 0
// 	dp[0][1] = -prices[0]

// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
// 	}

// 	return dp[n-1][0]
// }

//滚动数组
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	cash, hold := 0, -prices[0]

	precash, prehold := cash, hold

	for i := 1; i < len(prices); i++ {
		cash = max(precash,prehold + prices[i])
		hold = max(prehold,precash - prices[i])

		precash = cash
		prehold = hold
	}

	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
