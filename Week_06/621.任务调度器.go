/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
    tmp, lastCycle, maxFreq := make([]int, 26), 0, 0
    for _, v := range tasks {
        tmp[v - 'A']++
        if tmp[v - 'A'] > maxFreq { maxFreq = tmp[v - 'A'] }
    }
    for i := 0; i < 26; i++ {
        if maxFreq == tmp[i] { lastCycle++ }
    }
    return max((n + 1) * (maxFreq - 1) + lastCycle, len(tasks))
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
// @lc code=end

