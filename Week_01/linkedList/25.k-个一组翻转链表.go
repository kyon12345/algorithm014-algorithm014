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
func reverseKGroup(head *ListNode, k int) *ListNode {
	//递归
	curr := head
	count := 0

	//跳到第k+1个node
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}

	if count == k {
		//以k + 1个node为头的翻转后的链表
		curr = reverseKGroup(curr,k)
		//head 指向链表的头
		//curr 指向翻转后链表的头
		for ;count > 0;count-- {//开始翻转当前层的链表前半部分
			tmp := head.Next //tmp指向直链表的next
			head.Next = curr//head指向翻转的链表
			curr = head//替换翻转链表的头为head
			head = tmp//开始下一次循环,直链表的下一个值
		}
		head = curr
	}

	return head
}
// @lc code=end

