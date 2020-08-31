/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package main
// @lc code=start


//O(4^N/sqrt(n)) O(N)
//dfs
// func generateParenthesis(n int) []string {
// 	ret := make([]string, 0)

// 	generator(0, 0, n, "", &ret)

// 	return ret
// }

// func generator (l,r,max int,s string,res *[]string) {
// 	if l == max && r == max {
// 		*res = append(*res, s)
// 		return
// 	}
// 	if (l < max) {
// 		generator(l + 1, r, max, s + "(", res)
// 	}

// 	if (r < l) {
// 		generator(l, r + 1, max, s + ")", res)
// 	}
// }

//bfs
func generateParenthesis(n int) []string {
	ret := []string{}

	l := 0
	r := 0

	

}
// @lc code=end
