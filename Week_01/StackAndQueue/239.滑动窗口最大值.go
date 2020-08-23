/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

package main


func maxSlidingWindow(nums []int, k int) []int {
	//暴力解法,从第一个元素开始循环,每次循环求出后 k-1 个元素中的最大值
	//O(n*k) O(N - k + 1)
	// var ret []int

	// for i := 0; i < len(nums) - k + 1; i++ {
	// 	tmp := math.MinInt32
	// 	for j := i; j < i + k; j++ {
	// 		if nums[j] > tmp {
	// 			tmp = nums[j]
	// 		}
	// 	}
	// 	ret = append(ret,tmp)
	// }

	// return ret

	//双端队列 O(n) O(n)
	window := []int{}
	result := []int{}

	for i := 0; i < len(nums); i++ {
		if len(window) > 0 && window[0]	<= i - k {
			window = window[1:]
		}

		//挤压,只保留最大值
		for len(window) > 0 && nums[i] > nums[window[len(window) - 1]] {
			window = window[:len(window) - 1]
		}

		window = append(window,i)

		if i >= k -1 {
			result = append(result,nums[window[0]])
		}
	}

	return result

	//动态规划
}
// @lc code=end

