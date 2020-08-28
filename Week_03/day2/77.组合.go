/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start

//回溯+剪枝
//O(KCkn) O(Ckn)
// package cursion

func combine(n int, k int) [][]int {
	result := make([][]int,0)
	helper(1,n,k,[]int{},&result)
	return result
}


func helper(pointer,n,k int,current []int,result *[][]int) {
	if len(current) == k {
		*result = append(*result,append([]int{},current...))
		return
	}
	//pointer 0 - n
	for i := pointer;i <= n - (k - len(current)) + 1; i++ {
		helper(i + 1,n,k,append(current,i),result)
	}
}


// @lc code=end

