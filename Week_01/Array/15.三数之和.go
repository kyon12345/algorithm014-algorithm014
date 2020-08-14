/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	//排序,双指针
	sort.Ints(nums)

	ret := [][]int{}

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		target,j,k := -num,i + 1,len(nums) - 1

		for j < k {
			if nums[j] + nums[k] < target {
				j++
			} else if nums[j] + nums[k] > target {
				k--
			} else {
				ret = append(ret,[]int{num,nums[j],nums[k]})
				for ;j < k && nums[j]==nums[j + 1];j++ {}
				for ;j < k && nums[k]==nums[k - 1];k-- {}
				j,k = j + 1,k - 1
			}
		}
	}

	return ret
}

// @lc code=end

