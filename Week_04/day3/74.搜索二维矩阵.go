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

	height,width := len(matrix),len(matrix[0])

	lo,hi := 0, height * width - 1

	for lo < hi {
		mid := lo + (hi - lo) >> 1

		if matrix[mid/width][mid%width] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return matrix[lo/width][lo%width] == target
}
// @lc code=end

