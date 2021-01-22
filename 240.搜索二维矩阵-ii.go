/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

//从左下或者右上开始搜索
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return 0
	}

	rows,cols := len(matrix),len(matrix[0]

	row,col := rows,0

	for	 row >= 0 && col < cols {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			row --
		} else {
			col ++
		}
	}

	return false
}	

// @lc code=end
