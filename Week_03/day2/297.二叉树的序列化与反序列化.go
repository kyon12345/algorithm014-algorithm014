/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */
package main

import (
	"strconv"
	"strings"
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
//dfs
type Codec struct {
    l []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ret := make([]string,0)
	
	rserialize(root, &ret)

	return strings.Join(ret, ",")
}

func rserialize(root *TreeNode,ret *[]string) {
	if root == nil {
		*ret = append(*ret,"nil")
		return
	}

	*ret = append(*ret,strconv.Itoa(root.Val))
	rserialize(root.Left, ret)
	rserialize(root.Right, ret)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	this.l = strings.Split(data, ",")

	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == "nil" {
		this.l = this.l[1:]
		return nil
	}

	root := &TreeNode{Val: myAtoi(this.l[0])}

	this.l = this.l[1:]

	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()

	return root
}

func myAtoi (s string) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		panic(a)
	}
	return a
}

//bfs
// type Codec struct {
    
// }

// func Constructor() Codec {
//     return Codec{}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
//     if root == nil {
// 		return ""
// 	}

// 	queue := []*TreeNode{root}
// 	c := []string{strconv.Itoa(root.Val)}

// 	for len(queue) > 0 {
// 		l := len(queue)

// 		for i := 0; i < l; i++ {
// 			if queue[i].Left != nil {
// 				queue = append(queue,queue[i].Left)
// 			}
// 			if queue[i].Right != nil {
// 				queue = append(queue,queue[i].Right)
// 			}
// 			add(&c, queue[i].Left)
// 			add(&c, queue[i].Right)
// 		}

// 		queue = queue[1:]
// 	}

// 	res := strings.Join(c, ",")
// 	return res
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {    
// 	c := strings.Split(data, ",")
	
// 	if len(data) == 0 {
// 		return nil
// 	}

// 	t := &TreeNode{Val: myAtoi(c[0])}
// 	queue := []*TreeNode{t}
	
// 	i := 1
// 	for len(queue) > 0 {
// 		l := len(queue)	
// 		for j := 0; j < l; j++ {
// 			if c[i] == "nil" {
// 				queue[j].Left = nil
// 			} else {
// 				queue[j].Left = &TreeNode{Val: myAtoi(c[i])}
// 				queue = append(queue,queue[j].Left)
// 			}
// 			i ++
// 			if c[i] == "nil" {
// 				queue[j].Right = nil
// 			} else {
// 				queue[j].Right = &TreeNode{Val: myAtoi(c[i])}
// 				queue = append(queue, queue[j].Right)
// 			}
// 			i ++
// 		}
// 		queue = queue[1:]
// 	}

// 	return t
// }

// func add(c *[]string, node *TreeNode) {
//     if node == nil {
//         *c = append(*c, "nil")
//     } else {
//         *c = append(*c, strconv.Itoa(node.Val))
//     }
// }

// func myAtoi(s string) int {
//     a, err := strconv.Atoi(s)
//     if err != nil {
//         panic(a)
//     }
//     return a
// } 



/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end
