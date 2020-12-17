package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{0, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{2, nil}

	n := head
	length := 0
	for n != nil {
		n = n.Next
		length++
	}

	fmt.Println(length)
}
