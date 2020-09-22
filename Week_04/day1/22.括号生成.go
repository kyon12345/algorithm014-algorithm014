/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package main
// @lc code=start

//dfs
// func generateParenthesis(n int) []string {
// 	res := make([]string,0)

// 	helperGenerator(0,0,n,"",&res)

// 	return res
// }

// func helperGenerator(l,r,n int,s string,res *[]string) {
// 	if l == n && r == n {
// 		*res = append(*res,s)
// 		return
// 	}

// 	if l < n {
// 		helperGenerator(l + 1,r,n,s + "(",res)
// 	}

// 	if r < l {
// 		helperGenerator(l,r + 1,n,s + ")",res)
// 	}
// }

//bfs
//手动维护栈

// type PairNode struct {
// 	left int
// 	right int
// 	s string
// }

// func generateParenthesis(n int) []string {
// 	res := make([]string,0)

// 	if n == 0 {
// 		return res
// 	}

// 	stack := []*PairNode{ {0,0,""}}

// 	for len(stack) > 0 {
// 		top := stack[len(stack) - 1]

// 		stack = stack[:len(stack) - 1]

// 		if top.left == n && top.right == n {
// 			res = append(res,top.s)
// 		}

// 		if top.left < n {
// 			stack = append(stack,&PairNode{top.left + 1,top.right,top.s + "("})
// 		}

// 		if top.right < top.left {
// 			stack = append(stack,&PairNode{top.left,top.right + 1,top.s + ")"})
// 		}
// 	}

// 	return res
// }


func generateParenthesis(n int) []string {
	res := make([]string,0)

	dfsPairs(0, 0, n, "", &res)

	return res
}

func dfsPairs(l,r,n int,s string,res *[]string) {
	if l == n && r == n {
		*res = append(*res,s)
		return
	}

	if l < n {
		dfsPairs(l + 1,r,n,s + "(",res)
	}

	if r < l {
		dfsPairs(l, r + 1, n, s + ")", res)
	}
}
// @lc code=end

