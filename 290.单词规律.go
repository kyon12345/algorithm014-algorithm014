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
	arrS := strings.Fields(s)

	if len(arrS) != len(pattern) {
		return false
	}

	hash1 := make(map[string]byte)
	hash2 := make(map[byte]string)

	for i := 0; i < len(pattern); i++ {
		v, ok := hash1[arrS[i]]
		v2, ok2 := hash2[pattern[i]]

		if ok && v != pattern[i] || ok2 && v2 != arrS[i] {
			return false
		} else {
			hash1[arrS[i]] = pattern[i]
			hash2[pattern[i]] = arrS[i]
		}
	}

	return true
}

// @lc code=end
