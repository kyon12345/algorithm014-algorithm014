/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package main

// @lc code=start

//使用words初始化tire
//循环每一个格子,向四个方向进行深度优先搜素
//如果node的word里面有值,哪么返回值,结束这一层搜索,回溯状态,继续搜索

type TireNode struct {
	next [26]*TireNode
	word string
}

func (this *TireNode) Insert(word string) {
	root := this

	for _, c := range word {
		if root.next[c-'a'] == nil {
			root.next[c-'a'] = &TireNode{}
		}
		root = root.next[c-'a']
	}

	root.word = word
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return nil
	}

	res := []string{}

	//init tire
	root := &TireNode{}
	for _, w := range words {
		root.Insert(w)
	}

	//dfs
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfsFind(board, i, j, &res, root)
		}
	}

	return res
}

func dfsFind(board [][]byte, i, j int, res *[]string, node *TireNode) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#' || node.next[c-'a'] == nil {
		return
	}

	node = node.next[c-'a']
	board[i][j] = '#'

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}


	dfsFind(board, i+1, j, res, node)
	dfsFind(board, i-1, j, res, node)
	dfsFind(board, i, j+1, res, node)
	dfsFind(board, i, j-1, res, node)

	board[i][j] = c
}

// @lc code=end
