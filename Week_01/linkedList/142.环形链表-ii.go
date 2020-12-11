/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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
func detectCycle(head *ListNode) *ListNode {
	//哈希表 O(n) O(n)
	// visitedNode := make(map[*ListNode]int)

	// node := head 

	// for node != nil {
	// 	if _,exists := visitedNode[node];exists {
	// 		return node
	// 	}
	// 	visitedNode[node] = 1
	// 	node = node.Next
	// }

	// return nil

	//快慢指针 O(n) O(1)
	//a = c + (n - 1)(b + c)
	slow := head
	fast := head

	if head == nil || head.Next == nil {
		return nil
	}

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
// @lc code=end

