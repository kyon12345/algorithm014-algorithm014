/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
package main
// @lc code=start
func firstUniqChar(s string) int {
	var arr [26]int

	for i,k := range s {
		arr[k - 'a'] = i
	}

	for i,k := range s {
		if arr[k - 'a'] == i {
			return i
		} else {
			arr[k - 'a'] = -1
		}
	}

	return -1
}
// @lc code=end

