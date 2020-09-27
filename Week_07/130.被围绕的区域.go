/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */
package main

// @lc code=start
//只有O才能和O连接,所有不被包围的O都和边界上的O连接
//dfs
// var n,m int

// func solve(board [][]byte)  {
// 	if len(board) == 0 || len(board[0]) == 0 {
// 		return
// 	}

// 	n,m = len(board),len(board[0])
// 	//从四条边向中间搜索标记相邻的O
// 	for i := 0; i < n; i++ {
// 		dfs(board,i,0)
// 		dfs(board,i,m - 1)
// 	}

// 	for i := 1; i < m - 1; i++ {
// 		dfs(board,0,i)
// 		dfs(board,n - 1,i)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if board[i][j] == 'A' {
// 				board[i][j] = 'O'
// 			} else if board[i][j] == 'O' {
// 				board[i][j] = 'X'
// 			}
// 		}
// 	}

// }

// func dfs(board [][]byte,x,y int) {
// 	if x < 0 || x >= n || y >= m || y < 0 || board[x][y] != 'O' {
// 		return
// 	}

// 	board[x][y] = 'A'
// 	dfs(board,x + 1,y)
// 	dfs(board,x -1 ,y)
// 	dfs(board, x, y - 1)
// 	dfs(board,x,y + 1)
// }

//union find
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	row := len(board)
	col := len(board[0])

	//加一个虚拟节点
	dummyNode := row * col
	uf := NewUnionFind(dummyNode + 1)

	//转化为一维
	node := func(a, b int) int {
		return a*col + b
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				if i == 0 || j == 0 || i == row-1 || j == col-1 {
					uf.Union(node(i, j), dummyNode)
				}

				if i > 0 && board[i-1][j] == 'O' {
					uf.Union(node(i, j), node(i-1, j))
				}

				if j > 0 && board[i][j-1] == 'O' {
					uf.Union(node(i, j), node(i, j-1))
				}

				if i < row - 1 && board[i + 1][j] == 'O' {
					uf.Union(node(i,j), node(i + 1,j))
				}

				if j < col - 1 && board[i][j + 1] == 'O' {
					uf.Union(node(i,j), node(i,j + 1))
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if !uf.Is_connected(node(i, j), dummyNode) && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

//find topest parent node
func (u *UnionFind) Find(i int) int {
	root := i

	for u.parent[root] != root {
		root = u.parent[root]
	}

	for u.parent[i] != i {
		i, u.parent[i] = u.parent[i], root
	}

	return root
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)

	if pi != pj {
		u.parent[i] = pj
		u.count--
	}
}

func (u *UnionFind) Is_connected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

// @lc code=end
