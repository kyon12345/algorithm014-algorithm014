/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	lo,hi := 0,n - 1

	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	rot := lo
	lo,hi = 0,n - 1
	for lo <= hi {
		mid := (lo + hi) >> 1
		realmid := (rot + mid) % n

		if nums[realmid] == target {
			return realmid
		} else if nums[realmid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

// @lc code=end
