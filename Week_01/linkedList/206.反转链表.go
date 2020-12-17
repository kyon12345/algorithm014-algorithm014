/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	//迭代 O(N) O(1)
	// var prev *ListNode = nil 
	// curr := head

	// for curr != nil {
	// 	tmp := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = tmp 
	// }

	// return prev

	//简化一下 O(n) O(n)
	// var list *ListNode

	// for head != nil {
	// 	list = &ListNode {
	// 		Val:head.Val,Next:list,
	// 	}
	// 	head = head.Next
	// }
	
	// return list

	//迭代
	var next *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = next
		next = headb
		head = tmp
	}

	return next
}
// @lc code=end

