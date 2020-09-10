/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import "container/heap"

// @lc code=start
//使用优先队列
func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int,0)

	for _, n := range nums {
		seen[n] ++
	}

	q := &priorityQueue{}

	heap.Init(q)

	for val,cnt := range seen {
		heap.Push(q, element{val:val,cnt:cnt})

		if q.Len() > k {
			heap.Pop(q)
		}
	}

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res,heap.Pop(q).(element).val)
	}


	return res
}

type element struct {
	val int
	cnt int
}

type priorityQueue  []element

func (pq priorityQueue) Len() int {return len(pq)}

func (pq priorityQueue) Less(i,j int) bool {return pq[i].cnt < pq[j].cnt}

func (pq priorityQueue) Swap(i,j int) {pq[i],pq[j] = pq[j],pq[i]}

func (q *priorityQueue) Push(x interface{}) {*q = append(*q,x.(element))}

func (q *priorityQueue) Pop() interface{} {
	ret := (*q)[len(*q) - 1]
	
	*q = (*q)[:len(*q) - 1]

	return ret
}

// @lc code=end
