package main

func main() {
	lengthOfLIS([]int{10,9,2,5,3,7,101,18})
}

func lengthOfLIS(nums []int) int {
	//[10,9,2,5,3,7,101,18]
	if len(nums) < 1 {
		return 0
	}

	dp := make([]int,len(nums))

	res := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0;j < i && nums[j] < nums[i];j ++ {
			dp[i] = max(dp[i],dp[j] + 1)
		}
		res = max(res, dp[i])
	}

	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
} 