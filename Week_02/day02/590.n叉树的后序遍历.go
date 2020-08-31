/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
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
// func postorder(root *Node) []int {
// 	ret := []int{}

// 	if root == nil {
// 		return ret
// 	}

// 	for _, c := range root.Children {
// 		ret = append(ret,postorder(c)...)
// 	}

// 	ret = append(ret,root.Val)

// 	return ret
// }


//bfs
func postorder(root *Node) []int {
	ret := []int{}

	if root == nil {
		return ret
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret,root.Val)

		for _, c := range root.Children {
			stack = append(stack,c)
		}
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i],ret[len(ret) - 1 - i] = ret[len(ret) - 1 - i],ret[i]
	}

	return ret
}

// @lc code=end

