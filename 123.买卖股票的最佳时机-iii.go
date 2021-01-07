/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */
package main

// @lc code=start
// func maxProfit(prices []int) int {
// 	//T[i][2][0] = max(T[i - 1][2][0],T[i - 1][2][1] + prices[i])
// 	//T[i][2][1] = max(T[i - 1][2][1],T[i - 1][1][0] - prices[i])
// 	//T[i][1][0] = max(T[i - 1][1][0],T[i - 1][1][1] + prices[i])
// 	//T[i][1][1] = max(T[i - 1][1][1],T[i - 1][0][0] - prices[i]) = max(T[i - 1][1][1],-prices)

// 	if len(prices) < 2 {
// 		return 0
// 	}

// 	n := len(prices)

// 	dp := make([][3][2]int, n)

// 	dp[0][1][0] = 0
// 	dp[0][1][1] = -prices[0]
// 	dp[0][2][0] = 0
// 	dp[0][2][1] = -prices[0]

// 	for i := 1; i < n; i++ {
// 		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
// 		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
// 		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
// 		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
// 	}

// 	return dp[n-1][2][0]
// }

//优化空间复杂度
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := [5]int{}

	dp[0] = 0
	dp[1] = -prices[0]
	dp[3] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[1] = max(dp[1], dp[0]-prices[i])
		dp[2] = max(dp[2], dp[1]+prices[i])
		dp[3] = max(dp[3], dp[2]-prices[i])
		dp[4] = max(dp[4], dp[3]+prices[i])
	}

	return max(dp[2], dp[4])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
