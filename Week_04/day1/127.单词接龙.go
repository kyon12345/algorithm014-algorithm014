/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */
package main
// @lc code=start
//bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]struct{})
	for _, word := range wordList {
		words[word] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return 0
	}
	queue := []string{beginWord}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	length := 1
	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]
			for j := 0; j < len(lastWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					newWord := lastWord[:j] + string(k) + lastWord[j+1:]
					if _, ok := words[newWord]; !ok || newWord == lastWord {
						continue
					}
					if newWord == endWord {
						return length + 1
					}
					if _, ok := visited[newWord]; !ok {
						queue = append(queue, newWord)
						visited[newWord] = struct{}{}
					}
				}
			}
		}
		length++
	}

	return 0
}
// @lc code=end

