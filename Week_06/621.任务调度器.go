/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */
package main
// @lc code=start
func leastInterval(tasks []byte, n int) int {
    //填桶思路,优先考虑数量最多的任务
    tmp,lastCycle,maxFreq := make([]int,26),0,0

    for _, t := range tasks {
        tmp[t - 'A'] ++

        if tmp[t - 'A'] > maxFreq {
            maxFreq = tmp[t - 'A']
        }
    }

    for i := 0; i < 26; i++ {
        if maxFreq == tmp[i] {lastCycle ++}
    }

    return max(len(tasks),(maxFreq - 1) * (n + 1) + lastCycle)
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
// @lc code=end

