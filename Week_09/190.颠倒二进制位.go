/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */
package main
// @lc code=start
func reverseBits(num uint32) uint32 {
	var result uint32
	
	for i := 0; i < 32; i++ {
		result += num >> i << 31 >> i
	}

	return result
}
// @lc code=end

