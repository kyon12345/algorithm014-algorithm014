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

	cols = make(map[int]bool)
	pie = make(map[int]bool)
	na = make(map[int]bool)

	ans = make([][]string, 0)

	dfs(n,0,[]int{})
	
	return ans
}

func dfs(n, row int, path []int) {
	if row >= n {
		ans = append(ans, getBoard(n, path))
		return
	}

	for col := 0; col < n; col++ {
		if cols[col] || pie[col+row] || na[row-col] {
			continue
		}

		cols[col] = true
		pie[col+row] = true
		na[row-col] = true

		dfs(n,row + 1,append(path,col))

		cols[col] = false
		pie[col + row] = false
		na[row - col] = false
	}
}

func getBoard(n int, path []int) []string {

	res := make([]string, 0)

	for _, col := range path {
		res = append(res, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
	}

	return res
}

// @lc code=end
