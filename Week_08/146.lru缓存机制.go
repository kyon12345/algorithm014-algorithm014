/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */
package main

// @lc code=start
type LRUCache struct {
	cap        int
	m          map[int]*LinkNode
	head, tail *LinkNode
}

type LinkNode struct {
	pre, next  *LinkNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	tail, head := &LinkNode{nil, nil, 0, 0}, &LinkNode{nil, nil, 0, 0}

	tail.pre = head
	head.next = tail

	return LRUCache{capacity, make(map[int]*LinkNode), head, tail}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.MoveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	cache := this.m
	tail := this.tail.pre

	if node,ok := this.m[key];ok {
		node.value = value
		this.MoveToHead(node)
	} else {
		if this.cap == len(cache) {
			this.RemoveNode(tail)
			delete(cache, tail.key)
		}

		newNode := &LinkNode{tail,tail,key,value}

		this.AddNode(newNode)
		
		cache[key] = newNode
	}
}

func (this *LRUCache) AddNode(node *LinkNode) {
	this.head.next.pre = node
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
