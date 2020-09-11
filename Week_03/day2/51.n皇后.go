/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */
package main

import "strings"

// @lc code=start

var ans [][]string

var cols map[int]bool
var pie map[int]bool
var na map[int]bool


func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}

	ans = [][]string{}

	cols = make(map[int]bool,n)
	pie = make(map[int]bool,n)
	na = make(map[int]bool,n)

	dfs(0,n,[]int{})

	return ans
}

func dfs(row, n int,path []int) {
	if row >= n {
		ans = append(ans,getBoard(n, path))
	}

	for col := 0; col < n; col++ {
		if cols[col] || pie[row + col] || na[row - col] {
			continue
		}

		cols[col] = true
		pie[row + col] = true
		na[row - col] = true

		dfs(row + 1,n,append(path,col))

		cols[col] = false
		pie[row + col] = false
		na[row - col] = false
	}
}


func getBoard(n int,path []int) []string {
	
	res := make([]string,0)
	
	for _,col := range path {
		res = append(res, strings.Repeat(".", col) + "Q" + strings.Repeat(".", n - col - 1))
	}

	return res
}
// @lc code=end

