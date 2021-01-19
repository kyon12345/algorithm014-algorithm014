/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */
package main

// @lc code=start
//如果K > n/2 那么相当于k为正无穷的情况
func maxProfit(k int, prices []int) int {
	//dp[i][j] means maximum profit from at most i transactions using prices[0...j]
	//1.doing noting or buying dp[i][j] = dp[i][j - 1]
	//2.sell,must bought a stock on day t = 0 -> j - 1
	//example,choose to buy in the day t:
	//dp[i][j] = max(prices[j] - prices[t] + dp[i -1][t - 1])  t : 0 -> j - 1
	//base case dp[0][0] = 0
	n := len(prices)

	if n < 2 {
		return 0
	}

	if k >= n>>1 {
		return maxProfitkInfinity(prices)
	}

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i < k+1; i++ {
		tmpMax := -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], tmpMax+prices[j])
			tmpMax = max(tmpMax, dp[i-1][j-1]-prices[j])
		}
	}

	return dp[k][n-1]
}

func maxProfitkInfinity(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
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
