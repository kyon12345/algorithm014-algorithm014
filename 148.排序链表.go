/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//merge-sort buttom-up O(nlogn) o(1)
func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	dummy.Next = head

	n := 0
	for head != nil {
		head = head.Next
		n++
	}

	for step := 1; step < n; step <<= 1 {
		pre := dummy
		cur := pre.Next

		for cur != nil {
			left := cur
			right := split(cur, step)
			cur = split(right, step)
			pre = merge(pre, right, left)
		}
	}

	return dummy.Next
}

func split(head *ListNode, step int) *ListNode {
	if head == nil {
		return nil
	}

	for i := 1; head.Next != nil && i < step; i++ {
		head = head.Next
	}

	right := head.Next
	head.Next = nil

	return right
}

func merge(prev, right, left *ListNode) *ListNode {
	node := prev

	for right != nil && left != nil {
		if right.Val < left.Val {
			node.Next = right
			right = right.Next
		} else {
			node.Next = left
			left = left.Next
		}
		node = node.Next
	}

	if left != nil {
		node.Next = left
	} else if right != nil {
		node.Next = right
	}

	for node.Next != nil {
		node = node.Next
	}

	return node
}

// @lc code=end
