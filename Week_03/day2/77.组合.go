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
	if n == 0 {
		return nil
	}

	res := make([][]int,0)

	backtrackCombine(1, n, k, []int{}, &res)

	return res
}

func backtrackCombine(pointer, n, k int, curr []int, res *[][]int) {
	if len(curr) == k {
		*res = append(*res, curr)
		return
	}

	for i := pointer; i <= n - (k - len(curr)) + 1 ; i++ {
		backtrackCombine(i + 1, n, k, append(curr, i), res)
	}
}

// @lc code=end
