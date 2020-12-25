/*
 * @lc app=leetcode.cn id=979 lang=golang
 *
 * [979] 在二叉树中分配硬币
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
 * } */
func distributeCoins(root *TreeNode) int {
	res := 0

	if root.Left != nil {
		res += distributeCoins(root.Left)
		root.Val += root.Left.Val - 1
		res += abs(root.Left.Val - 1)
	}

	if root.Right != nil {
		res += distributeCoins(root.Right)
		root.Val += root.Right.Val - 1
		res += abs(root.Right.Val - 1)
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// @lc code=end
