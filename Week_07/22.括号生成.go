/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string,0)

	dfs(0,0,n,"",&res)

	return res
}

func dfs(l,r,n int,s string,res *[]string) {
	if l == n && r == n {
		*res = append(*res,s)
		return
	}

	if l < n {
		dfs(l + 1,r,n,s + "(",res)
	}

	if r < l {
		dfs(l,r + 1,n,s + ")",res)
	}
}
// @lc code=end

