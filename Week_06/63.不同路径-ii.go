/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */
package main
// @lc code=start
//O(M * N)
//O(M)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	width := len(obstacleGrid[0])

	dp := make([]int,width)

	dp[0] = 1
	for _, v := range obstacleGrid {
		for j := 0; j < width; j++ {
			if v[j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j - 1]
			}
		}
	}

	return dp[width - 1]
}
// @lc code=end

