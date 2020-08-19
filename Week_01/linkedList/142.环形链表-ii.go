/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

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
	slow := head
	fast := head

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next 
		fast = fast.Next.Next

		if(slow == fast) {
			break
		}
	}

	result := head

	//Floyd 算法,第二次相遇的节点必定是环的人口
	for result != slow {
		result = result.Next
		slow = slow.Next
	}

	return result
}
// @lc code=end

