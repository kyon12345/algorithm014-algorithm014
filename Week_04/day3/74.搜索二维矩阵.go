/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	//二维矩阵 row = index/width col = index%width
	//使用二分法,mid = matrix[num/2][width/2]
	//每次向中间逼近
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	length,width := len(matrix),len(matrix[0])

	if target < matrix[0][0] || target > matrix[length - 1][width - 1] {
		return false
	}

	l,r := 0,length * width

	for l <= r {
		mid := (l + r) /2

		row := mid/width
		col := mid % width

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
// @lc code=end

