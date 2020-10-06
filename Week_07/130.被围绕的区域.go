/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */
package main

import "golang.org/x/text/message/pipeline"

// @lc code=start
//只有O才能和O连接,所有不被包围的O都和边界上的O连接
//dfs
// var n,m int
// func solve(board [][]byte)  {

// 	if len(board) == 0 {
// 		return
// 	}

// 	var (
// 		rows = len(board)
// 		cols = len(board[0])
// 		dfs func (i,j int)
// 	)

// 	//从四条边向中间四个方向搜索
// 	dfs = func(i, j int) {
// 		if i < 0 || i >= rows || j  < 0 || j >= cols || board[i][j] != 'O' {
// 			return
// 		}

// 		board[i][j] = 'A'

// 		dfs(i - 1,j)
// 		dfs(i + 1,j)
// 		dfs(i,j - 1)
// 		dfs(i,j + 1)
// 	}

// 	for i := 0; i < rows; i++ {
// 		dfs(i,0)
// 		dfs(i,cols - 1)
// 	}

// 	for j := 1; j < cols - 1 ; j++ {
// 		dfs(0,j)
// 		dfs(rows - 1,j)
// 	}

// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if board[i][j] == 'A' {
// 				board[i][j] = 'O'
// 			} else if board[i][j] == 'O' {
// 				board[i][j] = 'X'
// 			}
// 		}
// 	}
// }

//union find
func solve(board [][]byte)  {
	if len(board) == 0 {
		return
	}

	var (
		rows = len(board)
		cols = len(board[0])
		dummyNodes = rows * cols
		node func(i,j int) int
	)

	uf := NewUnionFind(dummyNodes + 1)
	
	node = func(i, j int) int {
		return i * cols + j
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				if i == 0 || j == 0 || i == rows - 1 || j == cols - 1 {
					uf.Union(dummyNodes, node(i,j))
				}

				if i > 0 && board[i - 1][j] == 'O' {
					uf.Union(node(i,j), node(i -1,j))
				}

				if j > 0 && board[i][j - 1] == 'O' {
					uf.Union(node(i,j), node(i,j - 1))
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' && !uf.Is_Connected(node(i,j), dummyNodes) {
				board[i][j] = 'X'
			}
		}
	}
	
}

type UnionFind struct {
		count int
		parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int,n)	

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &UnionFind{
		count: n,
		parent: parent,
	}
}

func (u *UnionFind) Find(i int) int {
	root := i

	for u.parent[root] != root {
		root = u.parent[root]
	}

	for u.parent[i] != i {
		i,u.parent[i] = u.parent[i],root
	}

	return root
}

func (u *UnionFind) Union(i,j int) {
	pi := u.Find(i)
	pj := u.Find(j)

	if pi != pj {
		u.parent[pi] = pj
		u.count --
	}
}

func (u *UnionFind) Is_Connected(i,j int) bool {
	return u.Find(i) == u.Find(j)
}
// @lc code=end
