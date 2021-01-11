/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */
package main

// @lc code=start
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}

	cash, hold := 0, -prices[0]

	precash, prehold := cash, hold
	for i := 1; i < len(prices); i++ {
		cash = max(precash, prehold+prices[i]-fee)
		hold = max(prehold, precash-prices[i])

		prehold = hold
		precash = cash
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
