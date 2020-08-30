/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
 
//dfs
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}

// 	root.Left,root.Right = root.Right,root.Left

// 	invertTree(root.Left)
// 	invertTree(root.Right)

// 	return root
// }

//bfs
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	s := []*TreeNode{root}

	for len(s) > 0 {
		node := s[len(s) - 1]
		node.Left,node.Right = node.Right,node.Left

		s = s[:len(s) - 1]

		if node.Left != nil {
			s = append(s,node.Left)
		}

		if node.Right != nil {
			s = append(s,node.Right)
		}
	}

	return root
}
// @lc code=end

