/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */
package main
// @lc code=start
//dfs
//本质是搜索连通块个数的问题
//从连通块的任意一点走都能到达任意节点
// func findCircleNum(M [][]int) int {
// 	visited := make([]int,len(M))
// 	count := 0

// 	for i := 0; i < len(M); i++ {
// 		//只找没有被访问过的学生
// 		if visited[i] == 0 {
// 			dfsCircle(M, visited, i)
// 			count ++
// 		}
// 	}

// 	return count
// }

// func dfsCircle(M [][]int,visited []int,i int) {
// 	//终止条件就是全部遍历完
// 	for j := 0; j < len(M); j++ {
// 		if M[i][j] == 1 && visited[j] == 0 {
// 			visited[j] = 1
// 			dfsCircle(M,visited,j)
// 		}
// 	}
// }

//并查集
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}

	length := len(M)
	ufind := NewUnionFind(length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if M[i][j] == 1 {
				ufind.Union(i, j)
			}
		}
	}

	return ufind.count
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

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

// 路径压缩后 查询为O(1)
func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
// @lc code=end

