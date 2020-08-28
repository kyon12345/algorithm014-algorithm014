/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// package main

// import (
// 	"strconv"	"strings"
// )


// type TreeNode struct {
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }

type Codec struct {
    l []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return rserialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	l := strings.Split(data,",")

	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			this.l = append(this.l,l[i]) 
		}
	}
	
	return this.rdeserialize()
}

func rserialize(root *TreeNode,str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = rserialize(root.Left,str)
		str = rserialize(root.Right,str)
	}

	return str
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val,_ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val:val}
	this.l = this.l[1:]

	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()

	return root
}
/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

