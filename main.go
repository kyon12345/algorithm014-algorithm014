package main

import "fmt"

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	tails := make([]int, len(nums))

	size := 0
	for _, n := range nums {
		i, j := 0, size
		for i != j {
			m := (i + j) >> 1
			if n > tails[m] {
				i = m + 1
			} else {
				j = m
			}
		}

		tails[j] = n

		if i == size {
			size++
		}
	}
	
	fmt.Printf("%v",tails)

	return size
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	lengthOfLIS([]int{1,2,3,0,5})
}
