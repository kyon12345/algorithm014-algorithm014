/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	//state: s[i,j] 从i到j长的s是否是回文字符串
	//dp[i,j] = (s[i] == s[j]) && dp[i + 1,j - 1]
	//首字符和尾字符相等且去掉头尾它也是一个回文子串
	//初始化:单个字符肯定是回文 dp[i,i] = true
	//输出: if dp[i,j] == true记录字符长度
	//边界条件 [i + 1,j - 1]不构成区间 j - 1 - (i + 1) + 1 < 2 => j - i < 3
	//当字符串长度为0,1,2的时候为特殊情况
	l := len(s)


	//长度小于2的一定是回文字符串
	if l < 2 {
		return s
	}

	maxLength,begin := 1,0

	//初始化表格
	dp := make([][]bool,l)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, l)
	}

	charSlice := []byte(s)

	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if charSlice[i] != charSlice[j] {
				dp[i][j] = false
			} else if j - i < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i + 1][j - 1]
			}

			//找到最大值
			if dp[i][j] == true && j - i + 1 > maxLength {
				maxLength = j - i + 1
				begin = i
			}
		}
	}

	return string(charSlice[begin:begin + maxLength])
}
// @lc code=end

