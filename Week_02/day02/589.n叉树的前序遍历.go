/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

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
	ret := make([]int,0)

	if root == nil {
		return ret
	}

	stack := make([]*Node,0)

	stack = append(stack,root)

	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret,top.Val)

		n := len(top.Children)

		for i := n - 1; i >= 0; i-- {
			stack = append(stack,top.Children[i])
		}
	}

	return ret
}
// @lc code=end

