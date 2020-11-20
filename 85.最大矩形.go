/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */
package main
// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	//It took me a lot of time to accept and understand this solution, three more outputs I created to help myself understand it
// matrix
// 0 1 1 1 1 1 0
// 0 0 1 1 1 0 0
// 0 0 0 1 0 0 0

// height
// 0 1 1 1 1 1 0
// 0 0 2 2 2 0 0
// 0 0 0 3 0 0 0

// left
// 0 1 1 1 1 1 0
// 0 0 2 2 2 0 0
// 0 0 0 3 0 0 0

// right
// 7 6 6 6 6 6 7
// 7 7 5 5 5 7 7
// 7 7 7 4 7 7 7

// result
// 0 5 5 5 5 5 0
// 0 0 6 6 6 0 0
// 0 0 0 3 0 0 0

// matrix
// 1 1 1 0 0 0 0 0
// 1 1 1 1 1 1 1 1
// 0 0 0 0 1 1 1 1
// 0 0 0 0 1 1 1 1

// height
// 1 1 1 0 0 0 0 0
// 2 2 2 1 1 1 1 1
// 0 0 0 0 2 2 2 2
// 0 0 0 0 3 3 3 3

// left
// 0 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 0
// 0 0 0 0 4 4 4 4
// 0 0 0 0 4 4 4 4

// right
// 3 3 3 8 8 8 8 8
// 3 3 3 8 8 8 8 8
// 8 8 8 8 8 8 8 8
// 8 8 8 8 8 8 8 8

// result
// 3 3 3 0 0 0 0 0
// 6 6 6 8 8 8 8 8
// 0 0 0 0 8 8 8 8
// 0 0 0 0 12 12 12 12

// matrix
// 0 0 1 0 0
// 0 0 1 0 0
// 0 0 1 0 0
// 0 1 1 1 0
// 0 1 1 1 0

// height
// 0 0 1 0 0
// 0 0 2 0 0
// 0 0 3 0 0
// 0 1 4 1 0
// 0 2 5 2 0

// left
// 0 0 2 0 0
// 0 0 2 0 0
// 0 0 2 0 0
// 0 1 2 1 0
// 0 1 2 1 0

// right
// 5 5 3 5 5
// 5 5 3 5 5
// 5 5 3 5 5
// 5 4 3 4 5
// 5 4 3 4 5

// results
// 0 0 1 0 0
// 0 0 2 0 0
// 0 0 3 0 0
// 0 3 4 3 0
// 0 6 5 6 0

// Note for myself: This last one in particular, i thought the result of row 4 col 2 may me "15" as the prev row will enable this row to "trust" the accumulated height, but is not like that, as right and left are getting inherited from the previous rows too. In other words, if you get a high height accumulated, you also get that column more restricted.
	if len(matrix) < 1 {
		return 0
	}

	m,n := len(matrix),len(matrix[0])

	right,left,height := make([]int,n),make([]int,n),make([]int,n)

	for i := 0; i < n; i++ {
		right[i] = n
		left[i] = 0
		height[i] = 0
	}

	maxA := 0
	for i := 0; i < m; i++ {
		cur_left,cur_right := 0,n

		//height
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j] ++
			} else {
				height[j] = 0
			}
		}

		//right edge
		for j := n - 1; j >= 0;j -- {
			if matrix[i][j] == '1' {
				right[j] = min(cur_right,right[j])
			} else {
				right[j] = n
				cur_right = j
			}
		}

		//left edge
		for j := 0;j < n;j ++ {
			if matrix[i][j] == '1' {
				left[j]	= max(cur_left, left[j])
			} else {
				left[j] = 0
				cur_left = j + 1
			}
		}

		//maxA
		for j := 0; j < n; j++ {
			maxA = max(maxA, (right[j] - left[j]) * height[j])
		}
	}

	return maxA
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}


func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}
// @lc code=end

