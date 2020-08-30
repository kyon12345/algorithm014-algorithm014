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

	ret := make([][]int,0)

	backtrack(nums,[]int{},&ret)
	return ret
}


func backtrack(nums []int,prev []int, ans *[][]int) {
	if len(nums) == 0 {
		// *ans = append(*ans,append([]int{},prev...))
		*ans = append(*ans,prev)
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrack(append(append([]int{},nums[:i]...),nums[i + 1:]...),
		append(prev,nums[i]),ans)
	}
}

// @lc code=end