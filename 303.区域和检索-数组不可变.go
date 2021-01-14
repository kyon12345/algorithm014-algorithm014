/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */
package main

// @lc code=start
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	return NumArray{sum: sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := this.sum
	return sum[j+1] - sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
