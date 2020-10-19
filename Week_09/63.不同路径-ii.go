/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */
package main
// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//if grid[i][j] == 1 dp[i][j] = 0
	//dp[i][j]代表走到第i行第j列的不同路径
	//dp[i][j] = dp[i - 1][j] + dp[i][j - 1]

	//init
	width := len(obstacleGrid[0])

	dp := make([]int,width)

	//base case
	dp[0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < width; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j - 1]
			}
		}
	}

	return dp[width - 1]
}
// @lc code=end

