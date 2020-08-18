/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package main

// @lc code=start
func trap(height []int) int {
	//暴力法
	//寻找左右边界
	//当前柱子雨滴的面积 = 最小的边界 - 当前高度
	// ans := 0
	// size := len(height)

	// for i := 1; i < size - 1;i ++ {
	// 	max_left,max_right := 0,0
	// 	for j := i ; j >= 0; j-- {
	// 		max_left = max(max_left,height[j])
	// 	}

	// 	for j := i; j < size; j++ {
	// 		max_right = max(max_right,height[j])
	// 	}

	// 	ans += min(max_left, max_right) - height[i];
	// }

	// return ans

	//动态编程
	//从左往右数一遍
	//从右往左数一遍
	//两者的最小就是面积
	// if height == nil {
	// 	return 0
	// }

	// ans := 0
	// size := len(height)

	// left_max, right_max := make([]int, size), make([]int, size)

	// left_max[0] = height[0]

	// for i := 1; i < size; i++ {
	// 	left_max[i] = max(height[i], left_max[i-1])
	// }

	// right_max[size-1] = height[size-1]

	// for i := size - 2; i >= 0; i-- {
	// 	right_max[i] = max(height[i], right_max[i+1])
	// }

	// for i := 0; i < size; i++ {
	// 	ans += min(left_max[i], right_max[i]) - height[i]
	// }

	// return ans

	//基于暴力法,优化时间复杂度
	//使用单调栈
	var stack []int
	ans := 0
	
	for i := 0; i < len(height); i++ {
		if len(stack) == 0 {
			stack = append(stack,i)
			continue
		}

		for len(stack) > 0 && height[i] > stack[len(stack) - 1] {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			ans += height[top] * (i - top - 1)
		}

		stack = append(stack,i)
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
