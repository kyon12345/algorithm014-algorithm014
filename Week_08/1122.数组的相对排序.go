/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */
package main
// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	//计数排序
	bucket := [1001]int{}

	for i := range arr1 {
		bucket[arr1[i]] ++
	}

	res := make([]int,len(arr1))

	j := 0
	for i := range arr2 {
		for bucket[arr2[i]] > 0 {
			res[j] = arr2[i]
			bucket[arr2[i]] --
			j ++
		}
	}

	for i := 0;i < 1001;i ++ {
		for bucket[i] > 0 {
			res[j] = i
			j ++
			bucket[i] --
		}
	}

	return res
}
// @lc code=end

