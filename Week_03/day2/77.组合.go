/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
package main

// @lc code=start

//回溯+剪枝
//O(KCkn) O(Ckn)
// package cursion

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	res := make([][]int,0)

	dfsCombine(1, n, k, []int{}, &res)

	return res
}

func dfsCombine(pointer,n,k int,prev []int, res *[][]int) {
	if len(prev) == k {
		*res = append(*res,append([]int{},prev...))
		return
	}

	for i := pointer; i <= n - (k - len(prev)) + 1; i++ {
		dfsCombine(i + 1, n, k, append(prev,i),res)
	}
}

// @lc code=end
