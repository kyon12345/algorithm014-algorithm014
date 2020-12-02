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

func (this *TrieNode) Insert(word string) {
	node := this

	for _, r := range word {
		r -= 'a'

		if node.children[r] == nil {
			node.children[r] = &TrieNode{}
		}

		node = node.children[r]
	}

	node.word = word
}

func findWords(board [][]byte, words []string) []string {
	trie := &TrieNode{}

	//init trie
	for _, w := range words {
		trie.Insert(w)
	}

	res := make([]string,0)

	//begin dfs
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, &res, trie)
		}
	}

	return res
}

func dfs(board [][]byte,i,j int,res *[]string,node *TrieNode) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#'	|| node.children[c - 'a'] == nil {
		return
	}

	node = node.children[c - 'a']

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	board[i][j] = '#'

	dfs(board, i + 1, j, res, node)
	dfs(board, i - 1, j, res, node)
	dfs(board, i, j + 1, res, node)
	dfs(board, i, j - 1, res, node)

	board[i][j] = c
}
// @lc code=end

