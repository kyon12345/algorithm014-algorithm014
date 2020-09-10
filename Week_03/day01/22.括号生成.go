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
type pairStr struct {
	left int
	right int
	s string
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	stack := []*pairStr{{0,0,""}}

	ret := []string{}

	for len(stack) > 0 {
		next := []*pairStr{}

		for _, node := range stack {
			if node.left < n {
				next = append(next,&pairStr{node.left + 1,node.right,node.s + "("})
			}

			if node.right < node.left {
				next = append(next,&pairStr{node.left,node.right + 1,node.s + ")"})
			}
		}

		for _, pair := range next {
			if pair.left == n && pair.right == n {
				ret = append(ret,pair.s)
			}
		}

		stack = next
	}

	return ret
}
// @lc code=end
