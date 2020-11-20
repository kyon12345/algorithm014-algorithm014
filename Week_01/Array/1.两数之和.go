/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main
// @lc code=start
func twoSum(nums []int, target int) []int {
	//1.暴力法 O(n^2) O(1) 
	// for i := 0; i < len(nums);i ++ {
	// 	for j := i + 1;j < len(nums);j ++ {
	// 		if(nums[i] + nums[j] == target) {
	// 			return []int{i,j}
	// 		}
	// 	}
	// }

	// return []int{}

	//2.哈希表 O(n) O(n)
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v,ok := m[target - nums[i]];ok {
			return []int{i,v}
		}
		m[nums[i]] = i
	}

	return []int{}
}
// @lc code=end

