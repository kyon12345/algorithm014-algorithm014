/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) (ans int) {
    ans = 0
    for i := 0; i < len(s); i++ {
        ans += check(i, i, s)
        ans += check(i, i + 1, s)
    }
    return
}

func check(i int, j int, s string) (cnt int) {
    cnt = 0
    for i >= 0 && j < len(s) && s[i] == s[j] {
        cnt += 1
        i -= 1
        j += 1
    }
    return
}
// @lc code=end

