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

	n := len(matrix)
	m := len(matrix[0])

	l,r := 0,m * n - 1

	for l != r {
		mid := (l + r) >> 1
		if matrix[mid / m][mid % m] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return matrix[r / m][r % m] == target
}
// @lc code=end

