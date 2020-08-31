/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

package main

 type Node struct {
	 Val int
	 Children []*Node
 }

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


// func preorder(root *Node) []int {
// 	ret := make([]int, 0)
//     var f func(*Node)

//     f = func(node *Node) {
//         if node == nil {
//             return 
//         }
//         ret = append(ret, node.Val)
//         for _, e := range node.Children {
//             f(e)
//         }
//     }
//     f(root)
//     return ret
// }

func preorder(root *Node) []int {

	if root == nil {
		return nil
	}

	stack := []*Node{root}
	ret := []int{}

	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		ret = append(ret,top.Val)

		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack,top.Children[i])
		}
	}

	return ret
}
// @lc code=end

