/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */
package main
// @lc code=start
type WordDictionary struct {
	next [26]*WordDictionary
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	node := this

	for _, v := range word {
		v -= 'a'

		if node.next[v] == nil {
			node.next[v] = &WordDictionary{}
		}

		node = node.next[v]
	}

	node.isEnd = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this

	for _, v := range word {
		if v == '.' {
			for j := 0; j < 26; j++ {
				if node.next[j] != nil {
					if node.next[j].Search(word[1:]) {
						return true
					}
				}
			}
		} else {
			if node.next[v - 'a'] == nil {
				return false
			}

			node = node.next[v - 'a']
		}
	}

	return node.isEnd
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

