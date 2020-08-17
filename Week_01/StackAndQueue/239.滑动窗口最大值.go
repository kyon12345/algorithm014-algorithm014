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
	result := make([]int, 0, len(nums)-k)
	window := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		//如果当前窗口已经不包含队列头的值,则弹出队列头
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		//如果新加到队列尾的值大于队列尾的值,则弹出队列尾
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		//添加当前的索引到队列尾
		window = append(window, i)
		//取出队列头的值到结果集
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result

	//动态规划
}
// @lc code=end

