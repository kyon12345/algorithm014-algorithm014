/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 */
package main

import (
	"math"
)


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

// 	res := make([]int,0)

// 	queue := []*TreeNode{root}
	

// 	for len(queue) > 0 {
// 		l := len(queue)
// 		max := math.MinInt64

// 		for i := 0; i < l; i++ {
// 			max = maxVal(max, queue[i].Val)

// 			if queue[i].Left != nil {
// 				queue = append(queue,queue[i].Left)
// 			}

// 			if queue[i].Right != nil {
// 				queue = append(queue,queue[i].Right)
// 			}
// 		}

// 		res = append(res,max)

// 		queue = queue[l:]
// 	}

// 	return res
// }

// func maxVal(a,b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

//dfs
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int,0)

	dfsLargestValuees(root, 0, &res)

	return res
}

func dfsLargestValuees(root *TreeNode,level int,res *[]int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res,math.MinInt64)
	}

	(*res)[level] = int(math.Max(float64((*res)[level]), float64(root.Val)))

	if root.Left != nil {
		dfsLargestValuees(root.Left, level + 1, res)
	}

	if root.Right != nil {
		dfsLargestValuees(root.Right, level + 1, res)	
	}
}
// @lc code=end
