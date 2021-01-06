/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */
package main

import (
	"image/jpeg"
	"math"
	"math/rand"
)

// @lc code=start
func kClosest(points [][]int, K int) [][]int {
	quickSelect(points, 0, len(points)-1, K)
	res := make([][]int, 0)
	for i := 0; i < K; i++ {
		res = append(res, points[i])
	}
	return res
}

func quickSelect(a [][]int, l, r, k int) {
	pivot := l
	for i := l; i < r; i++ {
		if compareDis(a[i], a[r]) {
			a[i], a[pivot] = a[pivot], a[i]
			pivot += 1
		}
	}

	a[pivot], a[r] = a[r], a[pivot]

	if pivot+1 == k {
		return
	} else if pivot+1 > k {
		quickSelect(a, l, pivot-1, k)
	} else {
		quickSelect(a, pivot+1, r, k)
	}
}

func compareDis(a []int, b []int) bool {
	return (a[0]*a[0] + a[1]*a[1]) <= (b[0]*b[0] + b[1]*b[1])
}

// @lc code=end
