/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
//递归
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	return maxVal(maxDepth(root.Left),maxDepth(root.Right)) + 1
// }

// func maxVal(a,b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

//广度优先搜素
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	ans := 0

	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue,node.Left)
			}

			if node.Right != nil {
				queue = append(queue,node.Right)
			}
			sz --
		}
		ans ++
	}

	return ans
}

// @lc code=end
