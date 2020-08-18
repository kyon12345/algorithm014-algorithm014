/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

package main

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	//双指针从前往后 O(n) O(n)
	//两个指针,分别指向nums1和nums2的第一个元素
	//如果num1的较大,则nums1的指针不动,移动nums2的指针
	//添加较小的元素到新的数组
	// i,j := 0,0
	// var nums3 []int
	// for i != m && j != n  {
	// 	if nums1[i] < nums2[j] {
	// 		nums3 = append(nums3,nums1[i])
	// 		i ++ 
	// 	} else if nums1[i] == nums2[j] {
	// 		nums3 = append(nums3,nums2[j])
	// 		nums3 = append(nums3,nums1[i])
	// 		j ++
	// 		i ++
	// 	} else {
	// 		nums3 = append(nums3,nums2[j])
	// 		j ++
	// 	}
	// }

	// if i == m {
	// 	nums3 = append(nums3,nums2[j:]...)
	// }

	// if j == n {
	// 	nums3 = append(nums3,nums1[i:]...)
	// }

	// copy(nums1, nums3)

	//简洁的写法
	//从后往前的双指针 O(n) O(1)
	for n > 0 {
		if m == 0 || nums2[n - 1] > nums1[m - 1] {
			nums1[m + n - 1] = nums2[n - 1]
			n --
		} else {
			nums1[m + n - 1] = nums1[m - 1]
			m --
		}
	}
}
// @lc code=end

