/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */
package main

// @lc code=start
func judgeCircle(moves string) bool {
	//模拟
	  x, y := 0, 0
    length := len(moves)
    for i := 0; i < length; i++ {
        switch moves[i] {
        case 'U':
            y--
        case 'D':
            y++
        case 'L':
            x--
        case 'R':
            x++
        }
    }
    return x == 0 && y == 0
}
// @lc code=end

