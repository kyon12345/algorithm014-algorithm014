/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package main

// @lc code=start
// func climbStairs(n int) int {
// 	p,q,r := 0,0,1

// 	for i := 0; i < n; i++ {
// 		p = q
// 		q = r
// 		r = p + q
// 	}

// 	return r
// }

//dp
func climbStairs(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// @lc code=end
