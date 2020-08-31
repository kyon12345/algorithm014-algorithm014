/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */
package main

type Node struct {
	Val int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//dfs 
//O(N) O(logN)/O(n)
// func levelOrder(root *Node) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	res := [][]int{}

// 	dfs(root,0,&res)

// 	return res
// }


// func dfs(root *Node,level int,res *[][]int) {
// 	if root == nil {
// 		return
// 	}
// 	if len(*res) == level {
// 		*res = append(*res,[]int{})
// 	}

// 	(*res)[level] = append((*res)[level],root.Val)

// 	for _, c := range root.Children {
// 		dfs(c,level + 1,res)
// 	}
// }


//bfs
//O(N) O(N)
func levelOrder(root *Node) [][]int {
	stack := []*Node{root}

	if root == nil {
		return nil
	}

	ret := [][]int{}
	i := 0

	for len(stack) > 0 {
		ret = append(ret,[]int{})

		next := []*Node{}

		for _, n := range stack {
			ret[i] = append(ret[i], n.Val)

			for _, c := range n.Children {
				next = append(next,c)
			}
		}

		i ++

		stack = next
	}

	return ret
}


// @lc code=end

