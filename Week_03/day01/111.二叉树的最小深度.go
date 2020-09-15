/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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

//BFS O(N) O(N)
// package main

// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	depth := 1
// 	nodes := []*TreeNode{root}

// 	//其实就是找到那个两个左右子树都是空的节点
// 	for {
// 		newNodes := []*TreeNode{}

// 		for _, n := range nodes {
// 			if n.Left == nil && n.Right == nil {
// 				return depth
// 			}
// 			if n.Left != nil {
// 				newNodes = append(newNodes,n.Left)
// 			}
// 			if n.Right != nil {
// 				newNodes = append(newNodes,n.Right)
// 			}
// 		}
// 		depth ++
// 		nodes = newNodes
// 	}
// 	return depth
// }


 //递归 DFS O(N) O(H)
// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	left,right := minDepth(root.Left),minDepth(root.Right)

// 	if left == 0 || right == 0 {
// 		return left + right + 1
// 	}

// 	return 1 + minVal(left,right)
// }

func minVal (a,b int) int {
	if a > b {
		return b
	}
	return a
}

//BFS O(N) O(N)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	dep := 0
	for len(queue) > 0 {
		dep ++
		l := len(queue)
		for i := 0; i < l; i++ {
			root = queue[i]
			if root.Left == nil && root.Right == nil {
				return dep
			}
			if root.Left != nil {
				queue = append(queue,root.Left)
			}
			if root.Right != nil {
				queue = append(queue,root.Right)
			}
		}
		queue = queue[l:]
	}

	return dep
}



// @lc code=end

