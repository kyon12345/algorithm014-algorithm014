/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */
 package main

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
// func preorderTraversal(root *TreeNode) []int {
// 	//递归
// 	ret := make([]int,0)

// 	_generator(root,&ret)

// 	return ret
// }


// func _generator(root *TreeNode,ret *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	*ret = append(*ret,root.Val)
// 	_generator(root.Left,ret)
// 	_generator(root.Right,ret)
// }


func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)

	ret := []int{}

	if root == nil {
		return nil
	}

	for len(stack) > 0 || root != nil {
		if root != nil {
			ret = append(ret,root.Val)
			stack = append(stack,root)
			root = root.Left
		} else {
			node := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			root = node.Right
		}
	}

	return ret
}
// @lc code=end

