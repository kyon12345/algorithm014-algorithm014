/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func postorderTraversal(root *TreeNode) []int {
// 	ret := make([]int,0)
// 	_generator(root,&ret)
// 	return ret
// }

// func _generator(root *TreeNode,ret *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	_generator(root.Left,ret)
// 	_generator(root.Right,ret)
// 	*ret = append(*ret,root.Val)
// }

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	type trackedNode struct {
		node *TreeNode
		visited bool
	}

	stack := []*trackedNode{{node:root}}
	res := []int{}

	for len(stack) > 0 {
		n := stack[len(stack) - 1]

		if n.visited {
			res = append(res,n.node.Val)
			stack = stack[:len(stack) - 1]
			continue
		}

		if n.node.Right != nil {
			stack = append(stack,&trackedNode{node:n.node.Right})
		}

		if n.node.Left != nil {
			stack = append(stack,&trackedNode{node:n.node.Left})
		}

		n.visited = true
	}

	return res
}

// @lc code=end
