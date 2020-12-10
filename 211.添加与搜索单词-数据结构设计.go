/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */
package main

// @lc code=start
type WordDictionary struct {
	node *TireNode
}

type TireNode struct {
	children [26]*TireNode
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	node := &TireNode{}

	return WordDictionary{node : node}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	node := this.node

	for _, v := range word {
		if node.children[v - 'a'] == nil {
			node.children[v - 'a'] = &TireNode{}
		}
		node = node.children[v - 'a']
	}

	node.isEnd = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(word,0,this.node)
}

func match(word string,index int,node *TireNode) bool {
	if len(word) == index {
		return node.isEnd
	}

	if word[index] != '.' {
		return node.children[word[index] - 'a'] != nil && match(word, index + 1, node.children[word[index] - 'a'])
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
