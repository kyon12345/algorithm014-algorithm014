/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
package main

import "math"
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root, nil, nil)
}

func helper(root,max,min *TreeNode) bool {
	if root == nil {
		return true
	}
	
	if max != nil && root.Val >= max.Val {
		return false
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	return helper(root.Left,root, min) &&  helper(root.Right, max, root)
}

//中序遍历
// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	prev := math.MinInt64
// 	stack := []*TreeNode{}

// 	for len(stack) > 0 || root != nil {
// 		if root != nil {
// 			stack = append(stack,root)
// 			root = root.Left
// 		} else {
// 			top := stack[len(stack) - 1]
// 			stack = stack[:len(stack) - 1] 
// 			if top.Val > prev {
// 				prev = top.Val
// 			} else {
// 				return false
// 			}
// 			root = top.Right
// 		}
// 	}

// 	return true
// }
// @lc code=end

