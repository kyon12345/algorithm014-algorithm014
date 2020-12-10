/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	//sliding window
	//O(N) O(sigama)
	// m := map[byte]int{}	
	// n := len(s)

	// rk,ans := -1,0

	// for i := 0; i < n; i++ {
	// 	if i != 0 {
	// 		delete(m,s[i - 1])
	// 	}

	// 	for rk + 1 < n && m[s[rk + 1]] == 0 {
	// 		m[s[rk + 1]] ++
	// 		rk ++
	// 	}

	// 	ans = max(ans,rk - i + 1)
	// }

	// return ans
	dic := [256]int{}

	for i := 0; i < len(dic); i++ {
		dic[i] = -1
	}

	maxlen,start := 0,-1

	for i := 0; i < len(s); i++ {
		if dic[s[i]] > start {
			start = dic[s[i]]
		}

		dic[s[i]] = i

		maxlen = max(maxlen,i - start)
	}

	return maxlen
}

func max (x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

