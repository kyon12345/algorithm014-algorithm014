/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n-1

	//find the point
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	rot := lo
	lo, hi = 0, n-1

	//case [4,5,7,0,1,2] => [4,5,7,0,1,2,4,5,7]
	//realmid = (rot + (hi + rot)) >> 1
	//= ((root + (hi + rot)) / 2)%n
	//= (root + hi / 2) %n
	//=(root + mid) % n
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
