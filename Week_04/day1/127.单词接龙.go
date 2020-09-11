/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */
package main
// @lc code=start
//bfs,标准模板
//求无向图最短路径问题
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//这种解法如果wordlist很大就会超时
	wdict := map[string][]int{}

	for id,w := range wordList {
		for i := range w {
			k := w[0:i] + "*" + w[i + 1:]
			if _,ok := wdict[k];!ok {
				wdict[k] = []int{}
			}
			wdict[k] = append(wdict[k],id)
		}
	}

	used,l := make([]bool,len(wordList)),1

	wordList = append(wordList,beginWord)

	q := []int{len(wordList) - 1}

	for len(q) > 0 {
		nextQ := []int{}

		l ++

		for _, qid := range q {
			w := wordList[qid]

			for i := range w {
				k := w[0:i] + "*" + w[i + 1:]
				for _, wid := range wdict[k] {
					if wordList[wid] == endWord {
						return l
					}
					if !used[wid] {
						used[wid] = true
						nextQ = append(nextQ,wid)
					}
				}
			}
		}
		q = nextQ
	}

	return 0
}
// @lc code=end

