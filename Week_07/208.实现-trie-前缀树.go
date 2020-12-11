/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */
package main

// @lc code=start
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this

	for _, c := range word {
		if node.next[c - 'a'] == nil {
			node.next[c - 'a'] = &Trie{}
		}

		node = node.next[c - 'a']
	}

	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	
	for _, c := range word {
		if node.next[c - 'a'] == nil {
			return false
		}

		node = node.next[c - 'a']
	}

	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	
	for _, c := range prefix {
		if node.next[c - 'a'] == nil {
			return false
		}

		node = node.next[c - 'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// @lc code=end
