/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */
package main
// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	lo,hi := 0,len(arr) - 1

	for lo < hi {
		mid := (lo + hi) >> 1

		if arr[mid] < arr[mid + 1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
// @lc code=end

