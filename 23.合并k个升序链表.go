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
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeK(lists, 0, len(lists)-1)
}

func mergeK(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1

	return MergeTwo(mergeK(lists, l, mid), mergeK(lists, mid+1, r))
}

func MergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = MergeTwo(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwo(l2.Next, l1)
		return l2
	}
}

// @lc code=end
