package main

import "fmt"

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return kSelect(nums, 0, len(nums)-1, k)
}

func kSelect(nums []int, left int, right int, k int) int {
	if left == right {
		return nums[left]
	}

	p := right
	l := left
	r := right - 1
	for l <= r {
		for l <= r && nums[l] >= nums[p] {
			l++
		}

		for l <= r && nums[r] < nums[p] {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[l], nums[p] = nums[p], nums[l]
	p = l

	if p > k-1 {
		return kSelect(nums, left, p-1, k)
	} else if p < k-1 {
		return kSelect(nums, p+1, right, k)
	}

	fmt.Println(nums)

	return nums[p]
}

func lengthOfLongestSubstring(s string) int {
	//sliding window
	m := [256]int{}

	for i, _ := range m {
		m[i]-- 
	}

	start, res := -1, 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] > start {
			start = i
		}

		m[s[i]] = i

		if i-start > res {
			res = i - start
		}
	}

	return res
}

func main() {
	lengthOfLongestSubstring(" ")
}
