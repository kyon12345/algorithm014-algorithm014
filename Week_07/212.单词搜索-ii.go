/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package main
// @lc code=start
type TireNode struct {
	word string
	children [26]*TireNode
}


func findWords(board [][]byte, words []string) []string {
	root := &TireNode{}
	
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c - 'a'] == nil {
				node.children[c - 'a'] = &TireNode{}
			}
			node = node.children[c - 'a']
		}
		node.word = w
	}

	result := make([]string,0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i,j,board,root,&result)
		}
	}

	return result
}

func dfs(i,j int, board [][]byte, node *TireNode,result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#' || node.children[c - 'a'] == nil {
		return
	}

	node = node.children[c -  'a']

	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	dfs(i + 1,j,board,node,result)
	dfs(i - 1,j,board,node,result)
	dfs(i,j + 1,board,node,result)
	dfs(i,j - 1,board,node,result)
	board[i][j] = c
}
// @lc code=end

