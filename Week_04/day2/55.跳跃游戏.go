/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package main
// @lc code=start
func canJump(nums []int) bool {
	l := len(nums) - 1

	for i := l; i >= 0; i-- {
		if nums[i] + i >= l {
			l = i
		}
	}

	return l <= 0
}
// @lc code=end

