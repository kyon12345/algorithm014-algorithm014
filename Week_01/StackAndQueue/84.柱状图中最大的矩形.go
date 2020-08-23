/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
package main

// @lc code=start
func largestRectangleArea(heights []int) int {
	//暴力求解 O(n^2)
	// length:=len(heights)
    // maxAns:=0
    // for i:=0;i<length;i++{
    //     curMin := heights[i]
    //     for j:=i;j<length;j++{
    //         if heights[j] < curMin {
    //             curMin=heights[j]
    //         }
    //         curV := curMin*(j-i+1)
    //         if maxAns < curV {
    //             maxAns = curV
    //         }
    //     }
    // }
    // return maxAns
 
    //使用栈 O(n) O(n)
	var stack []int

	stack = append(stack,-1)

	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) != 1 && heights[i] <= heights[stack[len(stack) - 1]] {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			maxArea = maxVal(maxArea,heights[top] * (i - stack[len(stack) - 1] - 1))
		}
		stack = append(stack,i)
	}

	for len(stack) != 1 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		maxArea = maxVal(maxArea,heights[top] * (len(heights) - stack[len(stack) - 1] - 1))
	}

	return maxArea
}

func maxVal(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
// @lc code=end

