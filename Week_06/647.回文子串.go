/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */
package main
// @lc code=start
func countSubstrings(s string) (ans int) {
    //dp[i][j] 代表字符从第i到第j个位置的子字串
    //dp[i][j]若要为回文字符,那么dp[i + 1][j - 1]必定是回文字符
    dp := make([][]int, len(s))

    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(s)) 
    }

    cnt := 0

    //设置矩阵的对角线为1
    for i := 0; i < len(dp); i++ {
        dp[i][i] = 1
        cnt ++
    }

    //开始遍历从row = 0 col = 1这个位置开始
    for col := 1; col < len(s); col++ {
        for row := 0;row < col;row ++ {
            //如果子字符的长度为2且两个都相等,则标记
            //如果字符长度大于2,则看左下角的值是不是1(去头去尾),且头尾相等,标记为1
            if row == col - 1 && s[row] == s[col] {
                cnt ++
                dp[row][col] = 1
            } else if dp[row + 1][col - 1] == 1 && s[col] == s[row] {
                cnt ++
                dp[row][col] = 1
            }
        } 
    }

    return cnt
    
}
// func countSubstrings(s string) (ans int) {
//     ans = 0
//     for i := 0; i < len(s); i++ {
//         ans += check(i, i, s)
//         ans += check(i, i + 1, s)
//     }
//     return
// }

// func check(i int, j int, s string) (cnt int) {
//     cnt = 0
//     for i >= 0 && j < len(s) && s[i] == s[j] {
//         cnt += 1
//         i -= 1
//         j += 1
//     }
//     return
// }
// @lc code=end

