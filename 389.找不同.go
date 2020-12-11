/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */
package main
// @lc code=start
func findTheDifference(s string, t string) byte {
	//暴力,使用额外的map,遍历两次 O(N) O(N)
	//排序 + 一次遍历
	//a^b^a = b
	//所有字符都进行异或剩下的就是多余的
	ans := t[len(t) - 1]

	for i := 0; i < len(s); i++ {
		ans ^= s[i]
		ans ^= t[i]
	}

	return ans
}
// @lc code=end

