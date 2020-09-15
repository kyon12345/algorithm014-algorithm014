/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package main

// @lc code=start
//回溯
//O(3^M * 4^N) O(M + N)
var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }

    combinations = []string{}

    backtrackLetter(digits, 0, "")

    return combinations
}

func backtrackLetter(digits string,index int,combine string) {
    if index == len(digits) {
        combinations = append(combinations,combine)
        return
    }

    letters := phoneMap[string(digits[index])]

    for i := 0; i < len(letters); i++ {
        backtrackLetter(digits, index + 1, combine + string(letters[i]))
    }
}
// @lc code=end

