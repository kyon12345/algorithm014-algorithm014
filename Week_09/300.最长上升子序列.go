/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */
package main

import "fmt"

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
func lengthOfLIS(nums []int) int {
	// q := []int{^int(^uint(0) >> 1)}

	// fmt.Printf("%b", q[0])

	// for _, i := range nums {
	// 	if i > q[len(q) - 1] {
	// 		q = append(q, i)
	// 	} else {
	// 		l,r := 0,len(q)
	// 		for l < r {
	// 			m := l + (r - l) / 2
	// 			if q[m] < i {
	// 				l = m + 1
	// 			} else {
	// 				r = m
	// 			}
	// 		}

	// 		q[l] = i
	// 	}
	// }

	// return len(q) - 1

	q := []int{^int(^uint(0) >> 1)}

	for _, i := range nums {
		if i > q[len(q) - 1] {
			q = append(q, i)
		} else {
			l,r := 0,len(q)

			for l < r {
				m := l + (r - l) / 2
				if q[m] < i {
					l = m + 1
				} else {
					r = m
				}		
			}

			q[l] = i	
		}
	}

	return len(q) - 1
}
// @lc code=end
