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
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	n := 0
	for head != nil {
		head = head.Next
		n ++
	}

	for step := 1; step < n; step <<= 1 {
		prev := dummy
		cur := dummy.Next

		for cur != nil {
			left := cur
			right := split(left,step)
			cur = split(right,step)
			prev = merge(right,left,prev)
		}
	}

	return dummy.Next
}

func split(head *ListNode,step int) *ListNode {
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

func merge(left,right,prev *ListNode) *ListNode	 {
	cur := prev

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}

		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	} else if right != nil {
		cur.Next = right
	}

	for cur.Next != nil {
		cur = cur.Next
	}

	return cur
}

// @lc code=end
