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

	n := head
	l := 0
	for n != nil {
		n = n.Next
		l++
	}

	for step := 1; step < l; step <<= 1 {
		pre := dummy
		cur := pre.Next
		for cur != nil {
			left := cur
			right := split(cur, step)
			cur = split(right, step)
			pre = merge(left, right, pre)
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

func merge(l, r, pre *ListNode) *ListNode {
	head := pre
	for l != nil && r != nil {
		if l.Val < r.Val {
			head.Next = l
			l = l.Next
		} else {
			head.Next = r
			r = r.Next
		}
		head = head.Next
	}

	if l != nil {
		head.Next = l
	} else if r != nil {
		head.Next = r
	}

	for head.Next != nil {
		head = head.Next
	}

	return head
}

// @lc code=end
