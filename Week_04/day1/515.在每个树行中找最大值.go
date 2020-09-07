/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 */
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
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

//bfs
// func largestValues(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	stack := []*TreeNode{root}
// 	res := []int{}

// 	for len(stack) > 0 {
// 		node := []*TreeNode{}
// 		max := math.MinInt64

// 		for _, n := range stack {
// 			if n.Val > max {
// 				max = n.Val
// 			}

// 			if n.Left != nil {
// 				node = append(node,n.Left)
// 			}

// 			if n.Right != nil {
// 				node = append(node,n.Right)
// 			}
// 		}

// 		res = append(res,max)

// 		stack = node
// 	}

// 	return res
// }

//dfs
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	dfsLargestValues(root, &res, 0)
	return res
}

func dfsLargestValues(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, math.MinInt32)
	}
	(*res)[level] = int(math.Max(float64((*res)[level]), float64(root.Val)))
	if root.Left != nil {
		dfsLargestValues(root.Left, res, level+1)
	}
	if root.Right != nil {
		dfsLargestValues(root.Right, res, level+1)
	}
}

// @lc code=end
