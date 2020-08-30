/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
//中序遍历
// func isValidBST(root *TreeNode) bool {
// 	var stack []*TreeNode
	
// 	pre := math.MinInt64
// 	for len(stack) > 0 || root != nil {
// 		for root != nil {
// 			stack = append(stack,root)
// 			root = root.Left
// 		}

// 		root = stack[len(stack) - 1]
// 		stack = stack[:len(stack) - 1]
// 		if root.Val <= pre {
// 			return false
// 		}

// 		pre = root.Val
// 		root = root.Right
// 	}

// 	return true
// }

//中序遍历递归
// var pre int

// func isValidBST(root *TreeNode) bool {
// 	pre = math.MinInt64
// 	return helper(root)
// }

// //中序遍历是递增的,这里的pre相当于哨兵节点
// func helper(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if !helper(root.Left) {
// 		return false
// 	}

// 	if root.Val <= pre {
// 		return false
// 	}

// 	pre = root.Val

// 	return helper(root.Right)
// }

//递归
// func isValidBST(root *TreeNode) bool {
// 	return helper(root,nil,nil)
// }

// func helper(root,min,max *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if max != nil && root.Val >= max.Val {
// 		return false
// 	}

// 	if min != nil && root.Val <= min.Val {
// 		return false
// 	}

// 	return helper(root.Left,min,root) && helper(root.Right,root,max)
// }


//中序遍历
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode

	pre := math.MinInt64

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if root.Val <= pre {
			return false
		}

		pre = root.Val
		root = root.Right
	}

	return true
}
// @lc code=end

