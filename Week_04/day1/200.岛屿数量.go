/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main
// @lc code=start
//dfs
//O(MN) O(MN)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}

			cnt ++

			helperIslands(grid,i,j)
		}
	}

	return cnt
}

func helperIslands(grid [][]byte,i,j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] ++

	helperIslands(grid, i + 1,j)
	helperIslands(grid, i - 1, j)
	helperIslands(grid, i, j + 1)
	helperIslands(grid, i, j - 1)
}
// @lc code=end

