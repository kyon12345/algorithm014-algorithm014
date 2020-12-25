/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */
package main

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	row := len(grid)

	if row == 0 {
		return 0
	}

	col := len(grid[0])

	maxArea := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			area := dfsIsland(grid, i, j)
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}

func dfsIsland(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j]++

	return 1 + dfsIsland(grid, i+1, j) + dfsIsland(grid, i-1, j) + dfsIsland(grid, i, j-1) + dfsIsland(grid, i, j+1)
}

// @lc code=end
