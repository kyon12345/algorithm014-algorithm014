/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
package main
// @lc code=start
func isAnagram(s string, t string) bool {

	//哈希表 O(n) O(1)
	a := [26]int{}
	b := [26]int{}

	for _, char := range s {
		a[char - 'a']++
	}

	for _, char := range t {
		b[char - 'a']++
	}

	return a == b
}
// @lc code=end

