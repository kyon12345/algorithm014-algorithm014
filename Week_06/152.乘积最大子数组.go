/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */
package main
// @lc code=start
// func maxProduct(nums []int) int {
// 	//fmax(i) = max{fmax(i - 1) * ai,fmin(i - 1) * ai,ai}
// 	//fmin(i) = min{fmax(i - 1) * ai,fmin(i - 1) * ai,ai}
// 	maxF,minF,ans := nums[0],nums[0],nums[0]

// 	for i := 1; i < len(nums); i++ {
// 		mx,mn := maxF,minF

// 		maxF = max(mx * nums[i],max(nums[i],mn * nums[i]))
// 		minF = min(mn * nums[i], min(nums[i],mx * nums[i]))
// 		ans = max(maxF,ans)
// 	}

// 	return ans
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }

func maxProduct(nums []int) int {
	tmp := nums[0]
	max := nums[0]
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max,min = min,max
		}

		max *= nums[i]
		if max < nums[i] {
			max = nums[i]
		}

		min *= nums[i]
		if min > nums[i] {
			min = nums[i]
		}

		if max > tmp {
			tmp = max
		}
	}

	return tmp
}
// @lc code=end

