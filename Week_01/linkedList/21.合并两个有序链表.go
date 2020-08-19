/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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
//     Val int
//     Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//暴力法
	//遍历l1比较l2的每个元素,如果比它小则往前排
	//如果等于则同时添加l1,l2
	//如果大于则添加l2,同时l2往后直到>=l1或为空
	//妹写出来


	//循环 O(n) O(n)
	// if l1 == nil {
	// 	return l2
	// } else if l2 == nil {
	// 	return l1
	// }

	// var l3First *ListNode = nil
	// var l3Last *ListNode = nil
	// var newNode *ListNode = nil

	// for true {
	// 	if l1.Val > l2.Val {
	// 		newNode =l2
	// 		l2 = l2.Next
	// 	} else {
	// 		newNode = l1
	// 		l1 = l1.Next
	// 	}

	// 	if l3Last == nil {
	// 		l3First = newNode
	// 		l3Last = newNode
	// 	} else {
	// 		l3Last.Next = newNode
	// 		l3Last = newNode
	// 	}

	// 	if l1 == nil {
	// 		l3Last.Next = l2
	// 		break
	// 	} else if l2 == nil {
	// 		l3Last.Next = l1
	// 		break
	// 	}
	// }

	// return l3First
	

	//递归 O(m + n) O(m + n)
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next,l1)
		return l2
	}
}
// @lc code=end

