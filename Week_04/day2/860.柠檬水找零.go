/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */
package main

// @lc code=start
func lemonadeChange(bills []int) bool {
	//因为只有三种可能所以穷举,优先使用高面值的钱币
	billsMap := map[int]int{5: 0, 10: 0}

	for _, v := range bills {
		if v == 5 {
			billsMap[5] ++
		}

		if v == 10 {
			if billsMap[5] > 0 {
				billsMap[10] ++
				billsMap[5] --
			} else {
				return false
			}
		}

		if v == 20 {
				if billsMap[10] > 0 && billsMap[5] > 0 {
					billsMap[10] --
					billsMap[5] --
				} else if billsMap[5] > 2 {
					billsMap[5] = billsMap[5] - 3
				} else {
					return false
				}
		}
	}

	return true
}

// @lc code=end
