/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package main
// @lc code=start
//递归 O(n * n!) O(n)
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := make([][]int,0)

	backtrackPermute(nums, []int{}, &res)

	return res
}

func backtrackPermute(nums []int,prev []int,res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res,append([]int{},prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrackPermute(append(append([]int{},nums[:i]...),nums[i + 1:]...), append(prev,nums[i]), res)
	}
}
// @lc code=end