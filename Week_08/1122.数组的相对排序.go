/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */
package main
// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	bucket := [1001]int{}

	for _, data := range arr1 {
		bucket[data] ++
	}

	res := []int{}
	for _, data := range arr2 {
		for bucket[data] > 0 {
			res = append(res, data)
			bucket[data] --
		}
	}

	for i := 0; i < 1001; i++ {
		for bucket[i] > 0 {
			res = append(res, i)
			bucket[i] --
		}
	}

	return res
}
// @lc code=end

