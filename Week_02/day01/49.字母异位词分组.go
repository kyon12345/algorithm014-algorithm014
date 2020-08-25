/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
package main

import "sort"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	//使用map + sort NKLogK O(NK)
	groups := make(map[string][]string)

	for _, str := range strs {
		b := []byte(str)

		sort.Slice(b,func(i,j int) bool {
			return b[i] < b[j]
		})

		key := string(b)

		groups[key] = append(groups[key],str)
	}

	ret := make([][]string,0,len(groups))

	for _, v := range groups {
		ret = append(ret,v)
	}

	return ret
}
// @lc code=end

