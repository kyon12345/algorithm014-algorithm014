/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */
package main
// @lc code=start
func sortArrayByParity(A []int) []int {
	for i,j := 0,0;j < len(A);j ++ {
		if A[j] & 1 == 0 {
			swap(A, i, j)
			i ++
		}
	}

	return A
}
// @lc code=end

