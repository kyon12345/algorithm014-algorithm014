/*
 * @lc app=leetcode.cn id=980 lang=golang
 *
 * [980] 不同路径 III
 */
package main

// @lc code=start
var res, empty int

func uniquePathsIII(grid [][]int) int {
	res = 0
	empty = 1

	var sx, sy int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sx, sy = i, j
			} else if grid[i][j] == 0 {
				empty++
			}
		}
	}

	dfsPath(grid, sx, sy)

	return res
}

func dfsPath(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	c := grid[i][j]

	if c < 0 {
		return
	}

	if c == 2 {
		if empty == 0 {
			res++
		}
		return
	}

	grid[i][j] = -2
	empty--

	dfsPath(grid, i+1, j)
	dfsPath(grid, i-1, j)
	dfsPath(grid, i, j+1)
	dfsPath(grid, i, j-1)

	grid[i][j] = c
	empty ++
}

// @lc code=end
