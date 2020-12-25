/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

//从左下或者右上开始搜索
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	col, row := len(matrix[0])-1, 0

	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}

	return false
}

// @lc code=end
