/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
package main

import "math"


 type TreeNode struct {
	 Val int
	 Left *TreeNode
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
//递归
// func isValidBST(root *TreeNode) bool {
// 	return helper(root, nil, nil)
// }

// func helper(root,max,min *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if max != nil && root.Val >= max.Val {
// 		return false
// 	}

// 	if min != nil && root.Val <= min.Val {
// 		return false
// 	}

// 	return helper(root.Left,root,min) && helper(root.Right, max, root)
// }

//中序遍历
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{root}

	pre := math.MinInt64

	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		if top.Val < pre {
			return false
		}

		stack = stack[:len(stack) - 1]

		if top.Right != nil {
			stack  = append(stack,root.Right)
		}

		if top.Left != nil {
			stack = append(stack,top.Right)
		}

		pre = top.Val
	}

	return true
}
// @lc code=end

