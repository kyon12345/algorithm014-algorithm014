/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */
package main
// @lc code=start
func isPowerOfTwo(n int) bool {
	return n > 0 && n & (n - 1) == 0
}
// @lc code=end

