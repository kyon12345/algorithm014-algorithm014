/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */
 package main 

 type ListNode struct {
	 Val int
	 Next *ListNode
 }

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	//双指针
	if head == nil || head.Next == nil {
		return false
	}

	slow,fast := head,head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
// @lc code=end

