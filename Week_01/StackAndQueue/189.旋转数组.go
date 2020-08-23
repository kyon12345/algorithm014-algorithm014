/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

package main


// @lc code=start
func rotate(nums []int, k int)  {
	//双端队列,走一步弹出队列头,加到队列尾部 O(N)
	// end := 0
	// for i := 0; i < k; i++ {
	// 	end = nums[len(nums) - 1]
	// 	nums = nums[:len(nums) - 1]
	// 	nums = append([]int{end},nums...)
	// }
	// fmt.Println(nums)

	//暴力 O(n * k) O(1)
	// var tmp,previous int
	// for i := 0; i < k; i++ {
	// 	previous = nums[len(nums) - 1]
	// 	for j := 0; j < len(nums); j++ {
	// 		tmp = nums[j]
	// 		nums[j] = previous
	// 		previous = tmp
	// 	}
	// }

	//使用额外的数组 O(n) O(n)
	// a := make([]int,len(nums))

	// for i := 0; i < len(nums); i++ {
	// 	a[(i + k) % len(nums)] = nums[i]
	// }

	// for i := 0; i < len(a); i++ {
	// 	nums[i] = a[i]
	// }

	//翻转所有元素,再翻转前k个和后l-k个元素
	//O(n) O(1)
	// reverse(nums)
	// reverse(nums[:k%len(nums)])
	// reverse(nums[k%len(nums):])

	//直接提取后面的到前面
	//O(n) O(1)
	// k = k % len(nums)
	// copy(nums, append(nums[len(nums) - k:],nums[:len(nums) - k]...))
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i],arr[len(arr) - i - 1] = arr[len(arr) - i - 1],arr[i]
	}
}
// @lc code=end

