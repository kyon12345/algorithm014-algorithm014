/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

import "math"

// @lc code=start
//dp[amount] = min(dp[amount - nums[i]] for each num) + 1
//bottom up
func coinChange(coins []int, amount int) int {
	// state: amount
	// dp(amount): fewest number of coins to make amount
	// choice: use each of the coins
	// base: dp(0) = 0

	// dp(amount) = min(dp(amount-coin), for each coin)+1

	//11 [1,2,5]

	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i-c]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

//top down
// func coinChange(coins []int, amount int) int {
//     if amount < 1 {
//         return 0
//     }

//     return dfs(coins,amount,make([]int, amount + 1))
// }

// func dfs(coins []int,remainder int,dp []int) int {
//     if remainder < 0 {
//         return -1
//     }

//     if remainder == 0 {
//         return 0
//     }

//     if dp[remainder] != 0 {
//         return dp[remainder]
//     }

//     minVal := math.MaxInt32

//     for _, c := range coins {
//         changeResult := dfs(coins,remainder - c,dp)

//         if changeResult >= 0 && changeResult < minVal {
//             minVal = 1 + changeResult
//         }
//     }

//     if minVal == math.MaxInt32 {
//         dp[remainder] = -1
//     } else {
//         dp[remainder] = minVal
//     }

//     return dp[remainder]
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
