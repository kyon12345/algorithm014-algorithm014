/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	maxCur, maxSoFar := 0, 0

	for i := 1; i < len(prices); i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = max(0, maxCur)
		maxSoFar = max(maxCur, maxSoFar)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
