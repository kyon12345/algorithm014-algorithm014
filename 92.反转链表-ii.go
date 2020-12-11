/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 *
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy
	for i := 0; i < m - 1; i++ {
		pre = pre.Next
	}

	begin := pre.Next
	then := begin.Next

	//inserting between pre and begin
	for i := 0; i < n-m; i++ {
		pre.Next = then
		begin.Next = then.Next
		then = then.Next
		then.Next = begin
	}

	return dummy.Next
}

// @lc code=end
