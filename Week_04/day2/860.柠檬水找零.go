/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */
package main
// @lc code=start
func lemonadeChange(bills []int) bool {
	billsMap := map[int]int {5 : 0 , 10 : 0 ,20 : 0}

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			billsMap[5] ++
		}

		if bills[i] == 10 {
			if billsMap[5] > 0 {
				billsMap[10] ++
				billsMap[5] --
			} else {
				return false
			}
		}

		if bills[i] == 20 {
			if billsMap[10] > 0 && billsMap[5] > 0 {
				billsMap[10] --
				billsMap[5] --
			} else if billsMap[5] > 2{
				billsMap[5] -= 3
			} else {
				return false
			}
		}
	}

	return true
}
// @lc code=end

