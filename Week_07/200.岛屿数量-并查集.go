/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main


// @lc code=start
//dfs
//O(MN) O(MN)
// func numIslands(grid [][]byte) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}

// 	cnt := 0
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {
// 			if grid[i][j] != '1' {
// 				continue
// 			}

// 			cnt ++

// 			dfsBoard(grid,i,j)
// 		}
// 	}

// 	return cnt
// }

// func dfsBoard(grid [][]byte,i,j int) {
// 	if i <0 || i >= len(grid) || j <0 || j >= len(grid[0]) || grid[i][j] != '1' {
// 		return
// 	}

// 	grid[i][j] ++

// 	dfsBoard(grid, i + 1, j)
// 	dfsBoard(grid, i - 1, j)
// 	dfsBoard(grid, i, j - 1)
// 	dfsBoard(grid, i, j + 1)
// }

//使用并查集
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var (
		rows = len(grid)
		cols = len(grid[0])
		node func(row,col int) int
		dummy = rows * cols
	)

	u := New(dummy + 1)

	node = func(row, col int) int {
		return row * cols + col
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				if i < rows - 1 && grid[i + 1][j] == '1' {
					u.Union(node(i,j), node(i + 1,j))
				}

				if j < cols - 1 && grid[i][j + 1] == '1' {
					u.Union(node(i,j), node(i,j + 1))
				}
			} else {
				u.Union(node(i,j), dummy)
			}
		}
	}

	return u.count - 1
}

type UnionFind struct {
	count int
	parent []int
}

func New(n int) *UnionFind {
	par := make([]int,n)

	for i := 0; i < n; i++ {
		par[i] = i
	}

	return &UnionFind{
		count: n,
		parent: par,
	}
}

func (this *UnionFind) Find(i int) int {
	root := i

	for this.parent[root] != root {
		root = this.parent[root]
	}

	for this.parent[i] != i {
		i,this.parent[i] = this.parent[i],root
	}

	return root
}

func (this *UnionFind) Union(i,j int) {
	pi := this.Find(i)
	pj := this.Find(j)

	if pi != pj {
		this.parent[pi] = pj
		this.count --
	}
}

// @lc code=end

