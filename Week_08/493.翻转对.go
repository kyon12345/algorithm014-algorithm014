/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */
package main

// @lc code=start
func reversePairs(nums []int) int {
	// 归并排序
	return mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(arr []int, l int, r int) int {
	if l >= r {
		return 0
	}

	tmp,i,k := make([]int,r - l + 1),l,0

	mid := l + (r - l) >> 1

	count := mergeSort(arr, l, mid) + mergeSort(arr, mid + 1, r)

	for j,idx := mid + 1,l;j <= r;j ++ {
		for ;i <= mid && arr[i] < arr[j];i ++ {
			tmp[k],k = arr[i],k + 1
		}

		tmp[k],k = arr[j],k + 1

		for idx <= mid && (1 + arr[idx]) >> 1 <= arr[j] {
			idx ++
		}

		count += mid - idx + 1
	}

	copy(arr[l + k:], arr[i : mid + 1])
	copy(arr[l:], tmp[:k])

	return count
}

// @lc code=end
