/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package main

// @lc code=start

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

//递归 log(n) log(n)
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}

	return quickMul(x,n)
}

func quickMul(x float64,n int) float64 {
	if n == 0 {
		return 1
	}

	//如果我要知道x^n次方我只需要知道x^n/2就行了
	y := quickMul(x, n >> 1)

	if n & 1 == 1 {
		return y * y * x
	} else {
		return y * y
	}
}
//迭代
//o(log n) o(1)
// func myPow(x float64, n int) float64 {
// 	if n < 0 {
// 		x = 1/x
// 		n = -n
// 	}

// 	res := 1.0
// 	for n > 0 {
// 		if n & 1 == 1 {
// 			res *= x
// 		}
	
// 		x,n = x*x,n >> 1
// 	}

// 	return res
// }

// @lc code=end
