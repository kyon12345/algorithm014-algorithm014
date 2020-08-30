/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	var candidate int
	count := 0
	for i:=0;i<len(nums);i++{
		if count == 0{
			candidate = nums[i]
			count++
		}else{
			if nums[i] == candidate{
				count++
			}else{
				count--
			}
		}
	}
	return candidate
}
// @lc code=end

