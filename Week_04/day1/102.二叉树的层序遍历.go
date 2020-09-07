/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

 //dfs
// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	res := [][]int{}

// 	dfs(root,0,&res)

// 	return res
// }

// func dfs(root *TreeNode,level int,res *[][]int) {
// 	if root == nil {
// 		return
// 	}

// 	if level == len(*res) {
// 		*res = append(*res,[]int{})
// 	}

// 	(*res)[level] = append((*res)[level],root.Val)

// 	dfs(root.Left,level + 1,res)

// 	dfs(root.Right,level + 1,res)
// }

//bfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}


	res := make([][]int,0)
	curr := []*TreeNode{root}


	i := 0
	for len(curr) > 0 {
		next := []*TreeNode{}

		res = append(res,[]int{})

		for _, n := range curr {
			res[i] = append(res[i],n.Val)

			if n.Left != nil {
				next = append(next,n.Left)
			}

			if n.Right != nil {
				next = append(next,n.Right)
			}
		}

		curr = next

		i ++
	}

	return res
}
// @lc code=end

