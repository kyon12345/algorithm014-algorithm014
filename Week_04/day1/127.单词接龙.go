/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */
package main
// @lc code=start
//bfs,标准模板
//求无向图最短路径问题
// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	//这种解法如果wordlist很大就会超时
// 	wdict := make(map[string][]int)

// 	for id, wd := range wordList {
// 		for i := range wd {
// 			k := wd[:i] + "*" + wd[i + 1:]
// 			if _,ok := wdict[k];!ok {
// 				wdict[k] = []int{}
// 			}
// 			wdict[k] = append(wdict[k],id)
// 		}
// 	}

// 	wordList = append(wordList,beginWord)

// 	q := []int{len(wordList) - 1}


// 	used,l := make([]bool,len(wordList)),1

// 	for len(q) > 0 {
// 		nextQ := []int{}

// 		l ++

// 		for _,i := range q {
// 			w := wordList[i]
			
// 			for i := range w {
// 				k := w[:i] + "*" + w[i + 1:]
// 				for _,wid := range wdict[k] {
// 					if wordList[wid] == endWord {
// 						return l
// 					}
// 					if !used[wid] {
// 						used[wid] = true
// 						nextQ = append(nextQ,wid)
// 					}
// 				}
// 			}
// 		}

// 		q = nextQ
// 	}

// 	return 0
// }

//双向BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	step := 0 
	startQueue := []string{beginWord}
	endQueue := []string{endWord}

	startUsed := make([]bool,len(wordList))
	endUsed := make([]bool,len(wordList))


	wordMap := make(map[string]int)
	for i,w := range wordList {
		wordMap[w] = i
	}

	if idx,ok := wordMap[endWord];!ok {
		return 0
	} else {
		endUsed[idx] = true
	}

	for len(startQueue) > 0 {
		step ++
		l := len(startQueue)

		for i := 0; i < l; i++ {
			chars := []byte(startQueue[i])

			for j := 0; j < len(chars); j++ {
				o := chars[j]
				for c := 'a'; c <= 'z'; c++ {
					chars[j] = byte(c)
					idx,ok := wordMap[string(chars)]

					if !ok || startUsed[idx] {
						continue
					}

					if endUsed[idx] {
						return step + 1
					} else {
						startQueue = append(startQueue,wordList[idx])
						startUsed[idx] = true
					}
				}
				chars[j] = o
			}
		}

		startQueue = startQueue[l:]

		if len(startQueue) > len(endQueue) {
			startQueue,endQueue = endQueue,startQueue
			startUsed,endUsed = endUsed,startUsed
		}
	}

	return 0
}
// @lc code=end

