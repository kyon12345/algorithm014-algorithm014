/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */
package main
// @lc code=start
//f[i,j] = min(f[i + 1,j] ,f[i + 1,j + 1]) + a[i,j]	

//递归
//O(2^N) 超时
// func minimumTotal(triangle [][]int) int {
// 	return dfsTriangle(triangle,0,0)
// }

// func dfsTriangle(triangle [][]int,i,j int) int {
// 	if i == len(triangle) {
// 		return 0
// 	}

// 	return min(dfsTriangle(triangle,i + 1,j),dfsTriangle(triangle,i + 1,j + 1)) + triangle[i][j]
// }

//递归 + 记忆搜索
// var memo [][]int

// func minimumTotal(triangle [][]int) int {
// 	memo = make([][]int,len(triangle))

// 	for i := 0; i < len(memo); i++ {
// 		memo[i] = make([]int, len(triangle))
// 	}

// 	return dfsTriangle(triangle, 0,0)
// }

// func dfsTriangle(triangle [][]int,i,j int) int {
// 	if i == len(triangle) {
// 		return 0
// 	}

// 	if memo[i][j] != 0 {
// 		return memo[i][j]
// 	}

// 	memo[i][j] = min(dfsTriangle(triangle, i + 1, j),dfsTriangle(triangle, i + 1, j + 1)) + triangle[i][j]

// 	return memo[i][j]
// }

//dp
// func minimumTotal(triangle [][]int) int {
// 	//dp[i][j] = min(dp[i + 1][j] + dp[i + 1][j + 1]) + triangle[i][j]
// 	if len(triangle) == 0 {
// 		return 0
// 	}

// 	if len(triangle) == 1 {
// 		return triangle[0][0]
// 	}

//buttom up
// 	for i := len(triangle) - 2; i >= 0; i-- {
// 		for j := 0; j < len(triangle[i]); j++ {
// 			triangle[i][j] += min(triangle[i + 1][j],triangle[i + 1][j + 1])
// 		}
// 	}

// 	return triangle[0][0]
// }

//dp space(n)
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := triangle[len(triangle) - 1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j],dp[j + 1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

