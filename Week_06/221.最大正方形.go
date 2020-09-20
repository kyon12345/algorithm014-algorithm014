/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
package main
// @lc code=start
func maximalSquare(matrix [][]byte) int {
	//dp[i][j] = min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1
	if len(matrix) == 0 {
		return 0
	}

	row := len(matrix)
	col := len(matrix[0])
	cur := make([]int,col)

	pre := 0
	size := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := cur[j]
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				cur[j] = int(matrix[i][j] - '0')
			} else {
				cur[j] = min(pre,min(cur[j],cur[j - 1])) + 1
			}

			if size < cur[j] {
				size = cur[j]
			}

			pre = tmp
		}
	}

	return size * size
}

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
// @lc code=end

