/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */
package main

// @lc code=start
type WordDictionary struct {
	root *TireNode
}

type TireNode struct {
	next [26]*TireNode
	item string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := &TireNode{}
	return WordDictionary{root}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root

	for _, c := range word {
		if node.next[c-'a'] == nil {
			node.next[c-'a'] = &TireNode{}
		}
		node = node.next[c-'a']
	}

	node.item = word
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(word, 0, this.root)
}

func match (word string, k int, node *TireNode) bool {
	if k == len(word) {
		return !(node.item == "")
	}

	if word[k] != '.' {
		return node.next[word[k] - 'a'] != nil && match(word, k + 1, node.next[word[k] - 'a']);
	} else {
		for i := 0; i < 26; i++ {
			if node.next[i]	!= nil {
				if match(word, k + 1, node.next[i]) {
					return true
				}
			}
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
