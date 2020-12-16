/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */
package main
// @lc code=start
type UnionFind struct {
	count int
	parent []int
}

func New(n int) *UnionFind {
	par := make([]int,n)

	for i := 0; i < n; i++ {
		par[i] = i
	}

	return &UnionFind{n,par}
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
	pi,pj := this.Find(i),this.Find(j)

	if pi != pj {
		this.parent[pi] = pj
		this.count --
	}
}


func accountsMerge(accounts [][]string) [][]string {
	emails := map[string][]int{}

	
}
// @lc code=end

