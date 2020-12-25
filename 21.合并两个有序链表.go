/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 *
 */
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	} else if l2 == nil {
// 		return l1
// 	} else if l1.Val < l2.Val {
// 		l1.Next = mergeTwoLists(l1.Next, l2)
// 		return l1
// 	} else {
// 		l2.Next = mergeTwoLists(l2.Next,l1)
// 		return l2
// 	}
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return preHead.Next
}

// @lc code=end
