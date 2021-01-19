/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */
package main

// @lc code=start
type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), 0

	if row > 0 {
		col = len(matrix[0])
	}

	sums := make([][]int, row+1)
	for i := range sums {
		sums[i] = make([]int, col+1)
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sums := this.sums
	return sums[row2+1][col2+1] - sums[row2+1][col1] - sums[row1][col2+1] + sums[row1][col1]	
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
