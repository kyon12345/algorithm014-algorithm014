/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	n := head
	length := 1
	for n.Next != nil {
		n = n.Next
		length ++
	}

	k = k % length

	if k == 0 {
		return head
	}

	n.Next = head
	n = head
	for i := 1; i < length - k; i++ {
		n = n.Next
	}

	head = n.Next
	n.Next = nil

	return head
}

// @lc code=end
