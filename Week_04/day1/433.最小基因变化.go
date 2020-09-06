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
	set := make(map[string]struct{})

	for _, v := range bank {
		set[v] =struct{}{}
	}

	queue := []string{start}

	i := 0 
	for len(queue) > 0 {
		i ++
		l := len(queue)

		for _, v := range queue {
			for v2, _ := range set {
				if isNear(v,v2) {
					if v2 == end {
						return i
					}
					queue = append(queue,v2)
					delete(set,v2)
				}
			}
		}
		queue = queue[l:]
	}
	return -1
}

func isNear (start,end string) bool {
	mut := 0 
	for i := 0; i < 8; i++ {
		if start[i] != end[i] {
			mut ++
		}
	}

	return mut == 1
}
	
// @lc code=end

