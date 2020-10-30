package main

import "fmt"

func main() {
	lengthOfLIS([]int{1,3,4,1})
}

func lengthOfLIS(nums []int) int {
	q := []int{^int(^uint(0) >> 1)}

	fmt.Printf("%b", q[0])

	for _, i := range nums {
		if i > q[len(q) - 1] {
			q = append(q, i)
		} else {
			l,r := 0,len(q)
			for l < r {
				m := l + (r - l) / 2
				if q[m] < i {
					l = m + 1
				} else {
					r = m
				}
			}

			q[l] = i
		}
	}

	return len(q) - 1
}