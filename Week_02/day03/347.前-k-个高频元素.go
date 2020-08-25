/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import (
	"container/heap"
)

// @lc code=start
//使用优先队列

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, n := range nums {
		seen[n]++
	}

	q := &priorityQueue{}
	heap.Init(q)
	for val, cnt := range seen {
		heap.Push(q, element{value: val, count: cnt})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]int, k)
	for i := 1; i <= k; i++ {
		v := heap.Pop(q)
		if s, ok := v.(element); ok {
			ans[k-i] = s.value
		}
	}
	return ans
}

type element struct {
	value int
	count int
}

type priorityQueue []element

func (q priorityQueue) Len() int           { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return q[i].count < q[j].count }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(element))
}

func (q *priorityQueue) Pop() interface{} {
	res := (*q)[len(*q) - 1]
	*q = (*q)[:len(*q) - 1]
	return res
}

// @lc code=end
