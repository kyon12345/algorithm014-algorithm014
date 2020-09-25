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

    for i := 0; i < len(tasks); i++ {
        tmp[tasks[i] - 'A'] ++
        if tmp[tasks[i] - 'A'] > maxFreq {
            maxFreq = tmp[tasks[i] - 'A']
        }
    }

    for i := 0; i < 26; i++ {
        if maxFreq == tmp[i] {lastCycle ++}
    }

    return max(len(tasks),(n + 1) * (maxFreq - 1) + lastCycle)
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
// @lc code=end

