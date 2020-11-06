/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
package main
// @lc code=start
func maximalSquare(matrix [][]byte) int {
	//dp[i][j] = min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1
	//当前位置取决于左上,上,左三个位置dp值的最小值,就是当前能得到的最小边长
	if len(matrix) < 1 {
		return 0
	}

	rows,cols := len(matrix),len(matrix[0])

	dp := make([]int,cols)

	pre,size := 0,0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tmp := dp[j]

			if matrix[i][j] == '0' || i == 0 || j == 0 {
				dp[j] = int(matrix[i][j] - '0')
			} else {
				dp[j] = min(tmp,min(dp[j - 1],pre)) + 1
			}

			pre = tmp

			if dp[j] > size {
				size = dp[j]
			}
		}
	}

	return size * size
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
// @lc code=end

