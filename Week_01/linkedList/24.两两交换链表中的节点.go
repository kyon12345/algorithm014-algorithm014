/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	//递归 O(n) O(n)
	// if head == nil || head.Next == nil {
	// 	return head
	// }

	// n := head.Next

	// head.Next = swapPairs(head.Next.Next)

	// n.Next = head

	// return n

	//迭代 O(N) O(1)
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next
		
		first.Next = second.Next
		second.Next = first
		pre.Next = second
		pre = first
	}

	return dummy.Next
}
// @lc code=end

