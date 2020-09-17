/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target > matrix[m - 1][n - 1] || target < matrix[0][0] {
		return false
	}

	left,right := 0,m*n - 1

	mid,r,c := 0,0,0

	for left <= right {
		mid = left + (right - left) >> 1
		r = mid/n
		c = mid%n

		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left >= m*n || matrix[left/n][left%n] != target {
		return false
	}

	return true
}
// @lc code=end

