/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
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

// type Node struct {
// 	Val int
// 	Children []*Node
// }
 //递归 O(n) o(logN)/O(n)

// var res [][]int

// func levelOrder(root *Node) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	res = [][]int{}

// 	dfs(root,0)
// 	return res

// }

// func dfs(root *Node,level int) {
// 	if root == nil {
// 		return
// 	}

// 	if len(res) == level {
// 		res = append(res,[]int{})
// 	}

// 	res[level] = append(res[level],root.Val)
// 	for _, n := range root.Children {
// 		dfs(n,level + 1)
// 	}
// }

//bfs 
// func levelOrder(root *Node) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	curr := []*Node{root}
// 	var res [][]int	

// 	i := 0 
// 	for len(curr) != 0 {
// 		res = append(res,[]int{})

// 		next := []*Node{}

// 		for _, c := range curr {
// 			res[i] = append(res[i],c.Val)
			
// 			for _, node := range c.Children {
// 				if node != nil {
// 					next = append(next,node)
// 				}
// 			}
// 		}

// 		curr = next
// 		i ++
// 	}
// 	return res
// }

// @lc code=end

