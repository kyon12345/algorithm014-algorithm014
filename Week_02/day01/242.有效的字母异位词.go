/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
package main

// @lc code=start
func isAnagram(s string, t string) bool {

	//哈希表 O(n) O(1)
	a, b := [26]int{}, [26]int{}

	for _, v := range s {
		a[v-'a']++
	}

	for _, v := range t {
		b[v-'a']++
	}

	return a == b
}

// @lc code=end
