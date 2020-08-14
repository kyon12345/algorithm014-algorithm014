/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	//动态规划 O(n) O(1)
	p,q,r := 0,0,1

	for i := 1; i <= n;i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}
// @lc code=end

