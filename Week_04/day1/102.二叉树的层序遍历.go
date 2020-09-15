/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */
package main

import "go.starlark.net/resolve"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//dfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int,0)

	dfs(root, 0, &res)

	return res
}

func dfs(root *TreeNode,level int,res *[][]int) {

	if root == nil {
		return
	}

	if level == len(*res) {
		*res = append(*res,[]int{})
	}

	(*res)[level] = append((*res)[level],root.Val)

	dfs(root.Left,level + 1,res)

	dfs(root.Right,level + 1,res)
}

//bfs
// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	res := make([][]int, 0)

// 	queue := []*TreeNode{root}

// 	for len(queue) > 0 {
// 		l := len(queue)

// 		q := []int{}

// 		for i := 0; i < l; i++ {
// 			q = append(q,queue[i].Val)
// 			if queue[i].Left != nil {
// 				queue = append(queue,queue[i].Left)
// 			}

// 			if queue[i].Right != nil {
// 				queue = append(queue,queue[i].Right)
// 			}
// 		}

// 		res = append(res,q)

// 		queue = queue[l:]
// 	}

// 	return res
// }

// @lc code=end
