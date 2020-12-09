/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package main
// @lc code=start

type TireNode struct {
	children [26]*TireNode
	word string
}

func (this *TireNode) Insert(word string) {
	node := this

	for _, v := range word {
		v -= 'a'

		if node.children[v] == nil {
			node.children[v] = &TireNode{}
		}

		node = node.children[v]
	}

	node.word = word
}

func findWords(board [][]byte, words []string) []string {
	//init tire
	node := &TireNode{}

	for _, w := range words {
		node.Insert(w)
	}

	res := make([]string,0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i,j,board,&res,node)
		}
	}

	return res
}

func dfs(i,j int,board [][]byte,res *[]string,node *TireNode) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	
	c := board[i][j]

	if c == '#' || node.children[c - 'a'] == nil {
		return
	}

	board[i][j] = '#'

	node = node.children[c - 'a']

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	dfs(i, j + 1, board, res, node)
	dfs(i, j - 1, board, res, node)
	dfs(i - 1, j, board, res, node)
	dfs(i + 1, j, board, res, node)


	board[i][j] = c
}

// @lc code=end

