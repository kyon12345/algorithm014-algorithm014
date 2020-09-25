/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package main
// @lc code=start
func uniquePaths(m int, n int) int {
	//dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	//只能从上走下来或者从左边走过来
	//使用滚动数组优化space为O(N)
	dp := make([]int,n)

	//base case
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j - 1]
			}
		}
	}

	return dp[n - 1]
}
// @lc code=end

