/*
 * @lc app=leetcode.cn id=282 lang=golang
 *
 * [282] 给表达式添加运算符
 */
package main

import "strconv"

// @lc code=start
func addOperators(num string, target int) []string {
	rst := []string{}

	if len(num) == 0 {
		return rst
	}

	helperAdd(&rst, "", num, target, 0, 0, 0)

	return rst
}

func helperAdd(rst *[]string, path, num string, target, pos, eval, multed int) {
	if pos == len(num) {
		if target == eval {
			*rst = append(*rst, path)
		}
		return
	}

	//1,2,3,6
	//("","123",6,0,0,0)
	
	for i := pos; i < len(num); i++ {
		if i != pos && num[pos] == '0' {
			break
		}
		cur, _ := strconv.Atoi(num[pos : i+1])
		if pos == 0 {
			helperAdd(rst, path+toString(cur), num, target, i+1, cur, cur)
		} else {
			helperAdd(rst, path+"+"+toString(cur), num, target, i+1, eval+cur, cur)
			helperAdd(rst, path+"-"+toString(cur), num, target, i+1, eval-cur, -cur)
			helperAdd(rst, path+"*"+toString(cur), num, target, i+1, eval-multed+multed*cur, multed*cur)
		}
	}
}

func toString(a int) string {
	return strconv.Itoa(a)
}

// @lc code=end
