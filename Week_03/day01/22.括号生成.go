/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

package main

func generateParenthesis(n int) []string {
	res := []string{}
	generator(0, 0, n, "", &res)
	return res
}

func generator(l int, r int, max int, s string, res *[]string) {
	if l == max && r == max {
		*res = append(*res, s)
		return
	}

	if l < max {
		generator(l+1, r, max, s+"(", res)
	}

	if r < l {
		generator(l, r+1, max, s+")", res)
	}
}

// @lc code=end
