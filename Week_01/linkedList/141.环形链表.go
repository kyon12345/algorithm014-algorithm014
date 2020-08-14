/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	//哈希表 O(n) O(n)
	// m := make(map[*ListNode]int)
	
	// for head != nil {
	// 	if _,exists := m[head];exists {
	// 		return true
	// 	}
	// 	m[head] = 1
	// 	head = head.Next
	// }

	// return false

	//双指针 O(n) O(1)
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

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

