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
	row := len(grid)
	if row == 0 {
		return 0
	}

	col := len(grid[0])
	//向右,向下
	dir := [][]int{{1,0},{0,1}}
	dummy_node := row * col

	//多开一个虚拟的空间
	uf := NewUnionFind(dummy_node + 1)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				uf.Union(i * col + j, dummy_node)
			}

			if grid[i][j] == '1' {
				for _, d := range dir {
					new_x := i + d[0]
					new_y := j + d[1]

					if new_x < row && new_y < col && grid[new_x][new_y] == '1' {
						uf.Union(i * col + j , new_x * col + new_y)
					}
				}
			}
		}
	}

	return uf.count - 1
}


type UnionFind struct {
	count int
	parent []int
}

func NewUnionFind(n int) *UnionFind	{
	parent := make([]int,n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count: n,
		parent: parent,
	}
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
// @lc code=end

