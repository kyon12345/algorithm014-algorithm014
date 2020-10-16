/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */
package main
// @lc code=start
type LRUCache struct {
	m map[int]*LinkNode
	head,tail *LinkNode
	cap int
}

type LinkNode struct {
	pre,next *LinkNode
	key,val int
}


func Constructor(capacity int) LRUCache {
	head := &LinkNode{nil,nil,0,0}
	tail := &LinkNode{nil,nil,0,0}

	head.next = tail
	tail.pre = head

	return LRUCache{make(map[int]*LinkNode),head,tail,capacity}
}


func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v,exists := cache[key];exists {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	cache := this.m
	tail := this.tail
	if v,exists := cache[key];exists {
		this.MoveToHead(v)
		v.val = value
	} else {
		node := &LinkNode{nil,nil,key,value}
		if this.cap == len(cache) {
			delete(cache, tail.pre.key)
			this.RemoveNode(tail.pre)
		}
		cache[key] = node
		this.AddNode(node)
	}
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
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

