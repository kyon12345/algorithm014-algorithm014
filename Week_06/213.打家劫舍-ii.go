/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
package main
// @lc code=start
func rob(nums []int) int {
	//去头去尾取最大
	// p1,q1 := 0,0

	// if len(nums) == 1 {
	// 	return nums[0]
	// }

	// for i := 1; i < len(nums); i++ {
	// 	tmp := p1
	// 	p1 = max(q1 + nums[i],p1)
	// 	q1 = tmp
	// }

	// p2,q2 := 0,0

	// for i := 0; i < len(nums) - 1; i++ {
	// 	tmp := p2
	// 	p2 = max(q2 + nums[i],p2)
	// 	q2 = tmp
	// }

	// return max(p1,p2)
	if len(nums) ==  1{
		return nums[0]
	}

	return max(robHelper(nums, 0, len(nums) - 1),robHelper(nums, 1, len(nums)))
}

func robHelper(nums []int,l,r int) int {
	p,q := 0,0
	for i := l; i < r; i++ {
		tmp := p
		p = max(q + nums[i],p)
		q = tmp
	}
	return p
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

