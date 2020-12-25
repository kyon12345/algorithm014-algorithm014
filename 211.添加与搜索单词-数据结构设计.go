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
	next  [26]*TireNode
	isEnd bool
}

func (this *TireNode) Insert(word string) {
	root := this

	for _, c := range word {
		if root.next[c-'a'] == nil {
			root.next[c-'a'] = &TireNode{}
		}
		root = root.next[c-'a']
	}

	root.isEnd = true
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := TireNode{}

	return WordDictionary{&root}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.root.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(this.root,0,word)
}

func match(root *TireNode,index int,word string) bool {
	if len(word) == index {
		return root.isEnd
	}

	c := word[index]

	if c != '.' {
		return root.next[c - 'a'] != nil && match(root.next[c - 'a'], index + 1, word)
	} else {
		for i := 0; i < 26; i++ {
			if root.next[i] != nil {
				if match(root.next[i], index + 1, word) {
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
