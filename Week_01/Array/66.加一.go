/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
 package main

// @lc code=start
func plusOne(digits []int) []int {
	//如果是9则进一
	var n int = len(digits)
	for i := n -1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	//得到新的切片[1,0,0....n个0]
	var a = make([]int,n + 1)
	a[0] = 1
	return a
}
// @lc code=end

