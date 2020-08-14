/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	//快慢指针 O(n) O(1)
	noneZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[noneZeroIndex] = nums[noneZeroIndex], nums[i]
			noneZeroIndex++
		}
	}
}

// @lc code=end

