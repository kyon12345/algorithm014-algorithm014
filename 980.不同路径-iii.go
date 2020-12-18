/*
 * @lc app=leetcode.cn id=980 lang=golang
 *
 * [980] 不同路径 III
 */
package main
// @lc code=start
var res,empty int
func uniquePathsIII(grid [][]int) int {
	res = 0
	empty = 1

	var sx,sy int
	
	m,n := len(grid),len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				empty ++		
			} else if grid[i][j] == 1 {
				sx = i
				sy = j
			}
		}
	}

	dfsRoute(grid,sx,sy)

	return res
}

func dfsRoute(grid [][]int,x,y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] < 0 {
		return
	}

	if grid[x][y] == 2 {
		if empty == 0 {
			res ++
		}
		return
	}

	grid[x][y] = -2
	empty --

	dfsRoute(grid,x + 1,y)
	dfsRoute(grid,x - 1,y)
	dfsRoute(grid,x,y - 1)
	dfsRoute(grid,x,y + 1)

	grid[x][y] = 0
	empty ++
}

// @lc code=end
