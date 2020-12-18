/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main
// @lc code=start
//dfs
//O(MN) O(MN)
func numIslands(grid [][]byte) int {\
	if len(grid) == 0 {
		return 0
	}

	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}

			cnt ++

			dfsIsland(grid, i, j)
		}
	}

	return cnt
}

func dfsIsland(grid [][]byte,x,y int) {
	if x < 0 || y < 0 || x == len(grid) || y == len(grid[0]) || grid[x][y] != '1' {
		return
	}

	grid[x][y] ++

	dfsIsland(grid, x - 1, y)
	dfsIsland(grid, x + 1, y)
	dfsIsland(grid, x, y + 1)	
	dfsIsland(grid, x, y - 1)
}
// @lc code=end

