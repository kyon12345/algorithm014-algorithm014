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
	res := make([]string,0)

	if len(board) == 0 {
		return res
	}

	//构造树
	tree := &TireNode{}
	for _, v := range words {
		tree.Insert(v)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board,i,j,tree,&res)
		}
	}

	return res
}

func dfs(board [][]byte,i,j int,node *TireNode,res *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#' || node.children[c - 'a'] == nil {
		return
	}

	node = node.children[c - 'a']

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	
	dfs(board,i + 1,j,node,res)
	dfs(board,i - 1,j,node,res)
	dfs(board,i,j - 1,node,res)
	dfs(board,i,j + 1,node,res)

	board[i][j] = c

}
// @lc code=end

