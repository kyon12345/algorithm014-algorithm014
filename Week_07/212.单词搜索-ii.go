/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package main

// @lc code=start

type TireNode struct {
	next [26]*TireNode
	word string
}

func (this *TireNode) Insert(word string) {
	node := this

	for _, c := range word {
		if node.next[c - 'a'] == nil {
			node.next[c - 'a'] = &TireNode{}
		}

		node = node.next[c - 'a']
	}

	node.word = word 
}

func findWords(board [][]byte, words []string) []string {
	root := &TireNode{}

	//init tire
	for _, c := range words {
		root.Insert(c)
	}

	res := make([]string,0)

	//begin dfs
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i,j,board,root,&res)
		}
	}

	return res
}

func dfs(i,j int,board [][]byte,node *TireNode,res *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#' || node.next[c - 'a'] == nil {
		return
	}

	node = node.next[c - 'a']
	board[i][j] = '#'

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	dfs(i + 1, j, board, node, res)
	dfs(i, j + 1, board, node, res)
	dfs(i - 1, j, board, node, res)
	dfs(i, j - 1, board, node, res)

	board[i][j] = c
}

// @lc code=end
