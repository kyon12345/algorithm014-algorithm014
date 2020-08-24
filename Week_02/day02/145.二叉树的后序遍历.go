/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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

	// //这个方法挺类似于上颜色的
	type trackTreeNode struct {
		root    *TreeNode
		visited bool
	}

	var (
		stack []*trackTreeNode
		res   []int
	)

	stack = append(stack, &trackTreeNode{root: root})

	for len(stack) != 0 {
		node := stack[len(stack) - 1]

		if node.visited {
			stack = stack[:len(stack) - 1]
			res = append(res,node.root.Val)
			continue
		}

		if node.root.Right != nil {
			stack = append(stack,&trackTreeNode{root: node.root.Right})
		}

		if node.root.Left != nil {
			stack = append(stack,&trackTreeNode{root: node.root.Left})
		}

		node.visited = true
	}

	return res
	
}

// @lc code=end
