/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
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
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[1]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, l, mid))
}

func mergeTwoLists(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	if left.Val < right.Val {
		left.Next = merge(left.Next, right)
		return left
	} else {
		right.Next = merge(right.Next, left)
		return right
	}

}

// @lc code=end
