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

			dfsIslands(i,j,grid)
		}
	}

	return cnt
}

func dfsIslands(i,j int,grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] ++

	dfsIslands(i+1, j, grid)
	dfsIslands(i, j + 1, grid)
	dfsIslands(i - 1, j, grid)
	dfsIslands(i, j - 1, grid)
}
// @lc code=end

