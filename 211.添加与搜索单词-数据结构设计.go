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
	children [26]*TireNode
	isEnd    bool
}

func (this *TireNode) Insert(word string) {
	node := this

	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &TireNode{}
		}

		node = node.children[c-'a']
	}

	node.isEnd = true
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	node := &TireNode{}

	return WordDictionary{root: node}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.root.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(word,0,this.root)
}

func match (word string,index int,node *TireNode) bool {
	if len(word) == index {
		return node.isEnd
	}

	c := word[index]

	if c != '.' {
		return node.children[c - 'a'] != nil && match(word, index + 1, node.children[c - 'a'])
	} else {
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				if match(word, index + 1, node.children[i]) {
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
