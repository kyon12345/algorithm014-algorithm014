/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// package main

// type ListNode struct {
// 	Val int
// 	Next *ListNode
// }

func reverseKGroup(head *ListNode, k int) *ListNode {
	//递归
	count := 0
	curr := head

	for count != k && curr != nil {
		curr = curr.Next
		count ++
	}

	if count == k {
		//排序后的链表头
		curr = reverseKGroup(curr,k)

		for ;count > 0;count -- {
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
		}
		head = curr
	}

	return head
}
// @lc code=end

