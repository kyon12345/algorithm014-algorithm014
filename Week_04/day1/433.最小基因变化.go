/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */
package main

// import "reflect"

// @lc code=start

//bfs
func minMutation(start string, end string, bank []string) int {
	bankSet := make(map[string]struct{})

	for i := 0; i < len(bank); i++ {
		bankSet[bank[i]] = struct{}{}
	}

	if _,ok :=  bankSet[end];!ok {
		return -1
	}

	queue := []string{start}

	step := 0
	for len(queue) > 0 {
		step ++ 

		l := len(queue)
		for _, v := range queue {
			for v2 := range bankSet {
				if isNear(v, v2) {
					if v2 == end {
						return step
					}
					queue = append(queue,v2)
					delete(bankSet,v2)
				}
			}
		}

		queue = queue[l:]
	}

	

	return -1
}

func isNear(start,end string) bool {
	cnt := 0 

	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			cnt ++
		}
	}

	return cnt == 1
}
	
// @lc code=end

