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

	temp := make([]int,r - l + 1)

	mid := (l + r) >> 1

	i,k := l,0

	count := 0

	count = mergeSort(arr, l, mid) + mergeSort(arr, mid + 1, r)

	for j,idx := mid + 1,l;j <= r;j ++ {
		for ;i <= mid && arr[i] < arr[j];i ++ {
			temp[k],k = arr[i],k + 1
		}

		temp[k],k = arr[j],k + 1

		for idx <= mid && (arr[idx] + 1) >> 1 <= arr[j] {
			idx ++
		}
		count += mid - idx + 1
	}

	copy(arr[l + k:], arr[i:mid + 1])
	copy(arr[l:], temp[:k])

	return count
}

// @lc code=end
