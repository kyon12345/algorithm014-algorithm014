/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// package main

//  type Node struct {
// 	 Val int
// 	 Children []*Node
//  }

//递归
// func postorder(root *Node) []int {
// 	// if root == nil {
//     //     return []int{}
//     // }

//     // results := []int{}
//     // for _, v := range root.Children {
//     //     results = append(results, postorder(v)...)
//     // }

//     // results = append(results, root.Val)

// 	// return results
// }

//迭代
func postorder(root *Node) []int {
	// ret := make([]int, 0)
	// if root == nil {
	// 	return ret
	// }

	// stack := make([]*Node, 0)

	// stack = append(stack, root)

	// for len(stack) > 0 {
	// 	root := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	ret = append(ret, root.Val)

	// 	for _, node := range root.Children {
	// 		stack = append(stack, node)
	// 	}
	// }

	// for i := len(ret)/2 - 1; i >= 0; i-- {
	// 	opp := len(ret) - 1 - i
	// 	ret[i], ret[opp] = ret[opp], ret[i]
	// }

	// return ret
	ret := make([]int,0)

	if root == nil {
		return nil
	}
	
	stack := make([]*Node,0)

	stack = append(stack,root)

	for len(stack) > 0 {
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret,root.Val)

		for _, v := range root.Children {
			stack = append(stack,v)
		}
	}

	// 翻转slice
	for i := 0; i < len(ret)/2; i++ {
		ret[i],ret[len(ret) - 1 - i] = ret[len(ret) - i - 1],ret[i]
	}


	return ret
}

// @lc code=end

