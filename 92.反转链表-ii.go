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

	start := pre.Next
	then := start.Next

	//inserting then between pre and start,ensure then is always next of start
	for i := 0; i < n - m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	// 1 - 2 -3 - 4 - 5 ; m=2; n =4 ---> pre = 1, start = 2, then = 3
    // dummy-> 1 -> 2 -> 3 -> 4 -> 5
	// first reversing : dummy->1 - 3 - 2 - 4 - 5; pre = 1, start = 2, then = 4
    // second reversing: dummy->1 - 4 - 3 - 2 - 5; pre = 1, start = 2, then = 5 (finish)

	return dummy.Next
}

// @lc code=end
