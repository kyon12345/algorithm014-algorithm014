/*
 * @lc app=leetcode.cn id=983 lang=golang
 *
 * [983] 最低票价
 */
package main

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	n := days[len(days)-1]
	dp := make([]int, n+1)
	M := make(map[int]bool)
	for _, v := range days {
		M[v] = true
	}

	for i := 1; i < len(dp); i++ {
		if !M[i] {
			dp[i] = dp[i-1]
			continue
		}

		one := dp[i-1] + costs[0]
		seven := dp[max(i-7, 0)] + costs[1]
		thirty := dp[max(i-30, 0)] + costs[2]
		dp[i] = min(one, seven, thirty)
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(arr ...int) int {
	a := arr[0]
	for _, v := range arr {
		if v < a {
			a = v
		}
	}
	return a
}

// @lc code=end
