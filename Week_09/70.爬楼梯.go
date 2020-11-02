/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package main
// @lc code=start
func climbStairs(n int) int {
	p,q,r := 0,0,1

	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}
// @lc code=end

