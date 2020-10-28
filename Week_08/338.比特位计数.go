/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	arr := make([]int,num + 1)

	for i := 1;i <= num;i ++ {
		arr[i] = arr[i & (i - 1)] + 1
	}

	return arr
}
// @lc code=end

