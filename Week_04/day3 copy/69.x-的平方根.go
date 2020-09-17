/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package main

import "math"

// @lc code=start

// //二分查找
// func mySqrt(x int) int {
//     l, r := 0,x

//     for l <= r {
//         mid := (l + r) >> 1
//         temp := mid * mid

//         if temp < x {
//             l = mid + 1
//         } else if temp > x {
//             r = mid - 1
//         } else {
//             return mid
//         }
//     }

//     return r
// }

//牛顿迭代法
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    C, x0 := float64(x), float64(x)
    for {
        xi := 0.5 * (x0 + C/x0)
        if math.Abs(x0 - xi) < 1e-7 {
            break
        }
        x0 = xi
    }
    return int(x0)
}

// @lc code=end

