/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

	type TrackedNode struct {
		node *TreeNode
		visited bool
	}

	ret := []int{}

	stack := []*TrackedNode{}

	stack = append(stack,&TrackedNode{node:root})

	for len(stack) > 0 {
		top := stack[len(stack) - 1]

		if top.visited {
			ret = append(ret,top.node.Val)
			stack = stack[:len(stack) - 1]
			continue
		}

		if top.node.Right != nil {
			stack = append(stack,&TrackedNode{node:top.node.Right})
		}

		if top.node.Left != nil {
			stack = append(stack,&TrackedNode{node:top.node.Left})
		}

		top.visited = true
	}

	return ret
}

// @lc code=end
