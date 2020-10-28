/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package main
// @lc code=start
type TrieNode struct {
	word string
	children [26]*TrieNode
}

func (this *TrieNode) Insert(s string) {
	node := this

	for _, v := range s {
		v -= 'a'

		if node.children[v] == nil {
			node.children[v] = &TrieNode{}
		}

		node = node.children[v]
	}

	node.word = s
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}

	if len(board) == 0 {
		return res
	}

	//init trie
	node := &TrieNode{}

	for _, w := range words {
		node.Insert(w)
	}

	//begin dfs
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i,j,&res,node,board)
		}
	}

	return res
}

func dfs(i,j int,res *[]string,node *TrieNode,board [][]byte) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#' || node.children[c - 'a'] == nil  {
		return
	}

	node = node.children[c - 'a']

	if node.word != "" {
		*res = append(*res,node.word)
		node.word = ""
	}

	board[i][j] = '#'

	dfs(i + 1,j,res,node,board)
	dfs(i - 1,j,res,node,board)
	dfs(i,j + 1,res,node,board)
	dfs(i,j - 1,res,node,board)

	board[i][j]	= c
}
// @lc code=end

