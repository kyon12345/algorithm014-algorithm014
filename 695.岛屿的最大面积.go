/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */
package main

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	max_area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				max_area = max(max_area, AreaOfIsland(grid, i, j))
			}
		}
	}

	return max_area
}

func AreaOfIsland(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + AreaOfIsland(grid, i+1, j) + AreaOfIsland(grid, i-1, j) + AreaOfIsland(grid, i, j+1) + AreaOfIsland(grid, i, j-1)
	}

	return 0
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
