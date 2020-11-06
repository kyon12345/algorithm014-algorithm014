/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */
package main

// @lc code=start
type LRUCache struct {
	m          map[int]*LinkNode
	head, tail *LinkNode
	cap        int
}

type LinkNode struct {
	key, value int
	next, pre  *LinkNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &LinkNode{0, 0, nil, nil}, &LinkNode{0, 0, nil, nil}

	head.next = tail
	tail.pre = head

	return LRUCache{make(map[int]*LinkNode), head, tail, capacity}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m

	if v, ok := cache[key]; ok {
		this.MoveToHead(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	cache := this.m
	tail := this.tail.pre

	if v, ok := cache[key]; ok {
		v.value = value
		this.MoveToHead(v)
	} else {
		if this.cap == len(cache) {
			delete(cache, tail.key)
			this.RemoveNode(tail)
		}

		node := &LinkNode{key, value, nil, nil}

		this.AddNode(node)
		cache[key] = node
	}
}

func (this *LRUCache) AddNode(node *LinkNode) {
	this.head.next.pre = node
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
