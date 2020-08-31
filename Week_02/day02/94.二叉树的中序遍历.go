/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
// package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }


// func inorderTraversal(root *TreeNode) []int {
// 	//递归
// 	ret := []int{}
// 	st(root, &ret)
// 	return ret
// }

// func st(root *TreeNode,ret *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	st(root.Left,ret)
// 	*ret = append(*ret,root.Val)
// 	st(root.Right,ret)
// }

//dfs
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	ret := make([]int,0)

// 	dfs(root,&ret)

// 	return ret
// }

// func dfs(root *TreeNode, ret *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	dfs(root.Left,ret)
// 	*ret = append(*ret,root.Val)
// 	dfs(root.Right,ret)
// }


//bfs
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode,0)

	ret := make([]int,0)

	if root == nil {
		return nil
	}

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack,root)
			root = root.Left
		} else {
			node := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			ret = append(ret,node.Val)
			
			root = node.Right
		}
	}

	return ret
}


// @lc code=end

