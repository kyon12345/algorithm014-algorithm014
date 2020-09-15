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

	for _, v := range bank {
		bankSet[v] = struct{}{}
	}

	if _,ok := bankSet[end];!ok {
		return -1
	}

	queue := []string{start}

	step := 0
	for len(queue) > 0 {
		step ++
		l := len(queue)

		for i := 0; i < l; i++ {
			for _, v := range bank {
				if !isNear(queue[i], v) {
					continue
				}

				if v == end {
					return step
				} else {
					queue = append(queue,v)
					delete(bankSet, v)
				}
			}
		}

		queue = queue[l:]
	}

	return -1

}

func isNear (a,b string) bool {
	cnt := 0 

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt ++
		}
	}

	return cnt == 1
}
	
// @lc code=end

