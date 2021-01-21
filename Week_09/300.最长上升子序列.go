/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */
package main

import (
	"crypto/x509"
	"fmt"
)

// @lc code=start
// func lengthOfLIS(nums []int) int {
//dp[i]	= max(dp[j]) + 1  0 <= j < i && nums[i] > nums[j]
//10,9,2,5,3,7,101

//special
// if len(nums) < 1 {
// 	return 0
// }

// dp := make([]int,len(nums))

// res := 0
// for j := 0; j < len(nums); j++ {
// 	dp[j] = 1
// 	for i := 0; i < j; i++ {
// 		if nums[i] < nums[j] {
// 			dp[j] = max(dp[j],dp[i] + 1)
// 		}
// 	}

// 	res = max(res,dp[j])
// }

// return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

//二分查找
//啪的一下,很快啊 nlogn
// func lengthOfLIS(nums []int) int {
// 	//[10,9,2,5,3,7,101,18]
// 	if len(nums) < 1 {
// 		return 0
// 	}

// 	q := []int{^int(^uint(0) >> 1)}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > q[len(q) - 1] {
// 			q = append(q, nums[i])
// 		} else {
// 			l,r := 0,len(q)
// 			for l < r {
// 				mid := l + (r - l) >> 1
// 				if q[mid] < nums[i] {
// 					l = mid + 1
// 				} else {
// 					r = mid
// 				}
// 			}

// 			q[l] = nums[i]
// 		}
// 	}

// 	return len(q) - 1
// }

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	tails := make([]int, len(nums))

	size := 0
	for _, n := range nums {
		i, j := 0, size
		for i != j {
			m := (i + j) >> 1
			if n > tails[m] {
				i = m + 1
			} else {
				j = m
			}
		}

		tails[j] = n

		if i == size {
			size++
		}
	}

	return size
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
