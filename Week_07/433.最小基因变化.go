/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */
package main

// @lc code=start
//单向BFS
// func minMutation(start string, end string, bank []string) int {
// 	bankSet := make(map[string]struct{})

// 	for _, v := range bank {
// 		bankSet[v] = struct{}{}
// 	}

// 	if _,ok := bankSet[end];!ok {
// 		return -1
// 	}

// 	geneQueue := []string{start}

// 	step := 0
// 	for len(geneQueue) > 0 {
// 		l := len(geneQueue)
// 		step ++

// 		for i := 0; i < l; i++ {
// 			for w := range bankSet {
// 				if !isNear(geneQueue[i],w) {
// 					continue
// 				}

// 				if end == w {
// 					return step
// 				}

// 				geneQueue = append(geneQueue, w)
// 				delete(bankSet, w)
// 			}
// 		}

// 		geneQueue = geneQueue[l:]
// 	}

// 	return -1
// }
//双向BFS
func minMutation(start string, end string, bank []string) int {
	//init map
	wordMap := make(map[string]int, len(bank))

	for i, word := range bank {
		wordMap[word] = i
	}

	startQueue := []string{start}
	endQueue := []string{end}

	startUsed := make([]bool, len(bank))
	endUsed := make([]bool, len(bank))

	if i, ok := wordMap[end]; !ok {
		return -1
	} else {
		endUsed[i] = true
	}

	//begin bfs
	step := 0
	for len(startQueue) > 0 {
		l := len(startQueue)
		step ++

		for i := 0; i < l; i++ {
			for word, idw := range wordMap {
				if !isNear(word, startQueue[i]) || startUsed[idw] {
					continue
				}

				if endUsed[idw] {
					return step
				}

				startQueue = append(startQueue, word)
				startUsed[idw] = true
			}
		}

		startQueue = startQueue[l:]

		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}

	return -1
}

func isNear(a, b string) bool {
	cnt := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}

	return cnt == 1
}

// @lc code=end
