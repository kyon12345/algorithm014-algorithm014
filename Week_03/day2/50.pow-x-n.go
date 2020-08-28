/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package main
// @lc code=start
//快速幂 + 递归 log(n) log(n)
// func myPow(x float64, n int) float64 {
// 	if n >= 0 {
// 		return quickMul(x,n)
// 	}

// 	return 1.0/quickMul(x,-n)
// }

// func quickMul(x float64,n int) float64 {
// 	if n == 0 {
// 		return 1
// 	}

// 	y := quickMul(x, n/2)
// 	if n%2 == 0 {
// 		return y*y
// 	}

// 	return y*y*x
// }

//快速幂 + 迭代
//O(log(n)) O(1)
// func myPow(x float64, n int) float64 {
// 	if n >= 0 {
// 		return quickMul(x,n)
// 	}
	
// 	return 1.0/quickMul(x,-n)
// } 

// func quickMul(x float64,N int) float64 {
// 	ans := 1.0

// 	x_contribute := x

// 	//在对N进行二进制拆分的同时得出答案
// 	for N > 0 {
// 		if N % 2 == 1{
// 			ans *= x_contribute
// 		}

// 		x_contribute *= x_contribute

// 		N /= 2
// 	}

// 	return ans
// }

//迭代 简洁写法
func myPow(x float64,n int) float64 {
	if n < 0 {
		x,n = 1/x,-n
	}

	result := 1.0 

	for n > 0 {
		//n 为奇数
		if n&1 == 1{
			result *= x
		}
		//n = n / 2
		x,n = x*x,n >> 1
	}

	return result
}
// @lc code=end

