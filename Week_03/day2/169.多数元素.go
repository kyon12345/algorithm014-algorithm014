/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package main

// @lc code=start
//投票算法 O(N) O(1)
func majorityElement(nums []int) int {
	count := 0 

	var candidate int

	for _, n := range nums {
		if count == 0 {
			candidate = n
			count ++
		} else if candidate == n {
			count ++
		} else {
			count --
		}
	}

	return candidate
}
// @lc code=end

