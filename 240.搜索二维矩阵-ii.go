/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

//从左下或者右上开始搜索
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)

	if m == 0 {
		return false
	}

	n := len(matrix[0])

	//left-bottom
	row, col := m-1, 0

	for row >= 0 && col < n {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			row--
		} else {
			col++
		}
	}

	return false
}

// @lc code=end
