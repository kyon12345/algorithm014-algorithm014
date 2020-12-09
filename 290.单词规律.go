/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */
package main

import "strings"

// @lc code=start
func wordPattern(pattern string, s string) bool {
	//哈希表
	str_arr := strings.Fields(s)

	if len(pattern) != len(str_arr) {
		return false
	}

	hash := make(map[byte]string)
	hash2 := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		v,ok := hash[pattern[i]]
		v2,ok2 := hash2[str_arr[i]]

		if ok && v != str_arr[i] || ok2 && v2 != pattern[i] {
			return false
		} else {
			hash[pattern[i]] = str_arr[i]	
			hash2[str_arr[i]] = pattern[i]
		}
	}

	return true
}

// @lc code=end
