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
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	var (
		rows = len(board)
		cols = len(board[0])
		dummyNodes = rows * cols
		node func(row,col int) int
	)

	node = func(row, col int) int {
		return row * cols + col
	}

	u := New(dummyNodes + 1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == rows - 1 || j == 0 || j == cols - 1 {
					u.Union(node(i,j), dummyNodes)
				}

				if i > 0 && board[i - 1][j] == 'O' {
					u.Union(node(i,j),node(i - 1,j))
				}

				if j > 0 && board[i][j - 1] == 'O' {
					u.Union(node(i,j), node(i,j - 1))
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !u.Is_Connected(node(i,j), dummyNodes) && board[i][j]	 == 'O' {
				board[i][j]	= 'X'
			}
		}
	}
}

type UF struct {
	parent []int
	count int
}

func New(n int) *UF {
	parent := make([]int,n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &UF{parent: parent,count: n}
}

func (this *UF) Find(i int) int {
	root := i

	for this.parent[root] != root {
		root = this.parent[root]
	}

	//compression
	for this.parent[i] != i {
		i,this.parent[i] = this.parent[i],root
	}

	return root
}

func (this *UF) Union(i,j int) {
	pi := this.Find(i)
	pj := this.Find(j)

	if pi != pj {
		this.parent[pi] = pj
		this.count --
	}
}

func (this *UF) Is_Connected (i,j int) bool {
	return this.Find(i)	== this.Find(j)
}

// @lc code=end
