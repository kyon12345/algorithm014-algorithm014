/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	dfs(root, sum, &res, []int{})

	return res
}

func dfs(root *TreeNode, sum int, res *[][]int, prev []int) {
	if root.Left == nil && root.Right == nil && root.Val == sum {
		*res = append(*res, append(prev,root.Val))
		return
	}

	if root.Left != nil {
		dfs(root.Left, sum-root.Val, res, append(append([]int{}, prev...), root.Val))
	}

	if root.Right != nil {
		dfs(root.Right, sum-root.Val, res, append(append([]int{}, prev...), root.Val))
	}
}

// @lc code=end
