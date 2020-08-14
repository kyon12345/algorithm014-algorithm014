/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	//双指针 O(n) O(1) 
	l := 0
	r := len(height) - 1
	area := 0

	for(l < r) {
		temp := min(height[l],height[r])*(r - l)

		if temp > area {
			area = temp
		}

		if(height[l] > height[r]) {
			r --
		} else {
			l ++
		}
	}

	return area
}

func min (a int , b int) int {
	if(a > b) {
		return b
	}
	return a
}
// @lc code=end

