/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
//O(N) O(N)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	cnt := 0
	for _, v := range inorder {
		if v == preorder[0] {
			break
		}
		cnt ++
	}

	root.Left = buildTree(preorder[1:cnt + 1],inorder[:cnt])
	root.Right = buildTree(preorder[cnt + 1:], inorder[cnt + 1:])


	return root
}

// @lc code=end
