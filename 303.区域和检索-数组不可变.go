/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */
package main

// @lc code=start
//[1,2,3,4,5]
//[0,1,3,6,10,15]
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)

	for i := 1; i < len(nums)+1; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	return NumArray{sums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sums := this.sum
	return sums[j+1] - sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
