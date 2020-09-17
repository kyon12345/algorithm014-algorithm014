	/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
package main
// @lc code=start
// //O(N^2) O(1)
// func jump(nums []int) int {
// 	end := len(nums) - 1
// 	step := 0

// 	for end != 0 {
// 		for i := 0; i < end; i++ {
// 			if nums[i] >= end - i {
// 				end = i
// 				step ++
// 				break
// 			}
// 		}
// 	}

// 	return step
// }

//O(N) O(1)
func jump(nums []int) int {
	end := 0
	maxPosition := 0 
	step := 0

	for i := 0; i < len(nums) - 1; i++ {
		maxPosition = maxVal(maxPosition, nums[i] + i)

		if i == end {
			step ++
			end = maxPosition
		}
	}

	return step
}

func maxVal (a,b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

