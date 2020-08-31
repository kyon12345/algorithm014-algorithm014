package main

import (
	"container/heap"
)

//使用堆 O(nlogk)
func getLeastNumbers(arr []int, k int) []int {

	//1.直接排序 nlogn
	//2.快排 n
	//3.构造堆 nLogk
	if k == 0 || k > len(arr) {
		return []int{}
	}

	h := &intHeap{}
	heap.Init(h)

	//维护一个大小为k的最小堆,当arr的数比heap要小的时候弹出最大的元素,加入最小的元素
	for _, v := range arr {
		if h.Len() < k{
			heap.Push(h,v)
		} else {
			if (*h)[0] > v {
				heap.Pop(h)
				heap.Push(h,v)
			}
		}
	}

	res := []int{}
	for h.Len() > 0 {
		res = append(res,heap.Pop(h).(int))
	}

	return res
}

type intHeap []int

func (h intHeap) Len() int {
    return len(h)
}

func (h intHeap) Less(i,j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return res
}