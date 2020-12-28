/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */
package main

// @lc code=start
func kClosest(points [][]int, K int) [][]int {
	return quickSelect(points, 0, len(points)-1, K-1)
}

func quickSelect(points [][]int, lo, hi, idx int) [][]int {
	if lo > hi {
		return [][]int{}
	}

	j := partition(points, lo, hi)
}

func partition(points [][]int, lo, hi int) {
	v := points[lo]
	dist := v[0]*v[0] + v[1]*v[1]
	i, j := lo, hi+1

	for {
		
	}
}

// @lc code=end
