/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */
package main

// @lc code=start
type UnionFind struct {
	count  int
	parent []int
}

func New(n int) *UnionFind {
	par := make([]int, n)

	for i := 0; i < n; i++ {
		par[i] = i
	}

	return &UnionFind{
		count:  n,
		parent: par,
	}
}

func (this *UnionFind) Find(i int) int {
	root := i

	for this.parent[root] != root {
		root = this.parent[root]
	}

	for this.parent[i] != i {
		i, this.parent[i] = this.parent[i], root
	}

	return root
}

func (this *UnionFind) Union(i, j int) {
	pi := this.Find(i)
	pj := this.Find(j)

	if pi != pj {
		this.parent[pi] = pj
		this.count--
	}
}

func equationsPossible(equations []string) bool {
	uf := New(26)

	for _, c := range equations {
		if c[1] == '=' {
			uf.Union(int(c[0]) - 'a' , int(c[3]) - 'a')
		}
	}

	for _, c := range equations {
		if c[1] == '!' && uf.Find(int(c[0]) - 'a') == uf.Find(int(c[3] - 'a')) {
			return false
		}
	}

	return true
}

// @lc code=end
