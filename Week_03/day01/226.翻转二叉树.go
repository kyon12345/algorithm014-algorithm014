/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

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
// package main

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}

// 	root.Left,root.Right = root.Right,root.Left

// 	invertTree(root.Left)
// 	invertTree(root.Right)

// 	return root
// }

//BFS
// package main

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		var s []*TreeNode
		s = append(s,root)

		for len(s) > 0 {
			n := s[len(s) - 1]
			s = s[:len(s) - 1]

			n.Right,n.Left = n.Left,n.Right

			if n.Left != nil {
				s = append(s,n.Left)
			}

			if n.Right != nil {
				s = append(s,n.Right)
			}
		}
	}

	return root
}
// @lc code=end

