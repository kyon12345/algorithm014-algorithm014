/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */
package main

// @lc code=start
func sortArrayByParityII(A []int) []int {
	i,j,n := 0,1,len(A)

	for i < n && j < n {
		for i < n && A[i] & 1 == 0 {
			i += 2
		}

		for j < n && A[j] & 1 != 0 {
			j += 2
		}

		if i < n && j < n {
			swap(A, i, j)
		}
	}

	return A
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// @lc code=end
