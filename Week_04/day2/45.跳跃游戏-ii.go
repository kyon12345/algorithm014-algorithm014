	/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
package main
// @lc code=start
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 1

	currentJump,nextJump := nums[0],nums[0]	

	for i := 1; i < len(nums); i++ {
		if i == len(nums) - 1 {
			return jumps
		}

		if i + nums[i] > nextJump {
			nextJump = i + nums[i]
		}

		if i == currentJump {
			jumps ++
			currentJump = nextJump
		}
	}

	return jumps

}
// @lc code=end

