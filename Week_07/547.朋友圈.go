/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */
package main

import "text/template/parse"

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

	u := NewUF(len(M))

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if M[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}

	return u.count
}

type UnionFind struct {
	cap    int
	parent []int
}

func NewUF(cap int) *UnionFind {
	par := make([]int, cap)

	for i := 0; i < cap; i++ {
		par[i] = i
	}

	return &UnionFind{cap, par}
}

func (this *UnionFind) Find(index int) int {
	root := index

	for this.parent[root] != root {
		root = this.parent[root]
	}

	//copression
	for this.parent[index] != index {
		index,this.parent[index] = this.parent[index],root
	}

	return root
}

// @lc code=end
