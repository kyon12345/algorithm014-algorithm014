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
	seen := map[int]int{}

	for _, n := range nums {
		seen[n] ++
	}

	q := &priorityQueue{}
	heap.Init(q)

	for val,cnt := range seen {
		heap.Push(q,element{val,cnt})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]int,0)

	for i := 0; i < k; i++ {
		ans = append(ans,heap.Pop(q).(element).val)
	}

	return ans

}

type element struct {
	val int
	cnt int
}

type priorityQueue []element


func (q priorityQueue) Len() int {return len(q)}

func (q priorityQueue) Less(i,j int) bool {return q[i].cnt < q[j].cnt}

func (q priorityQueue) Swap(i,j int) {q[i],q[j] = q[j],q[i]}

func (q *priorityQueue) Push(x interface{}) {*q = append(*q,x.(element))}

func (q *priorityQueue) Pop() interface{} {
	ret := (*q)[q.Len() - 1]

	*q = (*q)[:q.Len() - 1]

	return ret
}
// @lc code=end
