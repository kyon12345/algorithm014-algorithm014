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
	res := [][]int{}

	if root == nil {
		return res
	}

	dfsSum(root, &res, []int{}, sum)

	return res
}

func dfsSum(root *TreeNode, res *[][]int, prev []int, sum int) {
	if root.Left == nil && root.Right == nil && root.Val == sum {
		prev = append(prev, root.Val)
		*res = append(*res, prev)
	}

	if root.Left != nil {
		dfsSum(root.Left, res, append(append([]int{}, prev...), root.Val), sum-root.Val)
	}

	if root.Right != nil {
		dfsSum(root.Right, res, append(append([]int{}, prev...), root.Val), sum-root.Val)
	}
}

// @lc code=end
