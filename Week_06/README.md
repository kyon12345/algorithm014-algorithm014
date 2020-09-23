学习笔记

DP:

-   分治子问题
-   状态数组定义
-   DP 方程

## 递归代码模板

func dfs() {
//terrimal condition

    //process result

    //process current logic

    //drill down

    //restore curren status

}

## 分治代码模板

func divide_conquer() {
//terrimal

    //prepare data

    //conquer subproblem

    //process and generate the final result

    //revert the current level status

}

## DP 的本质 就是分治 + 最优子结构

-   对于二维矩阵的问题,我们要考虑边界条件,有没有超出矩阵的范围
-   记忆状态数组来优化复杂度(递归)
-   所谓最优子结构,在简单的 DP 问题中一般是和,最小值,最大值
-   在某些情况下可以通过递归 + 记忆化搜索的方法来实现接近 DP 的时间复杂度

    ```
    //递归
      //O(2^N) 超时
      // func minimumTotal(triangle [][]int) int {
      // 	return dfsTriangle(triangle,0,0)
      // }

      // func dfsTriangle(triangle [][]int,i,j int) int {
      // 	if i == len(triangle) {
      // 		return **0**
      // 	}

      // 	return min(dfsTriangle(triangle,i + 1,j),dfsTriangle(triangle,i + 1,j + 1)) + triangle[i][j]
      // }

      func min(a,b int) int {
          if a < b {
              return a
          }
          return b
      }

      //递归 + 记忆搜索
      // var memo [][]int

      // func minimumTotal(triangle [][]int) int {
      // 	memo = make([][]int,len(triangle))

      // 	for i := 0; i < len(memo); i++ {
      // 		memo[i] = make([]int, len(triangle))
      // 	}

      // 	return dfsTriangle(triangle, 0,0)
      // }

      // func dfsTriangle(triangle [][]int,i,j int) int {
      // 	if i == len(triangle) {
      // 		return 0
      // 	}

      // 	if memo[i][j] != 0 {
      // 		return memo[i][j]
      // 	}

      // 	memo[i][j] = min(dfsTriangle(triangle, i + 1, j),dfsTriangle(triangle, i + 1, j + 1)) + triangle[i][j]

      // 	return memo[i][j]
      // }

    ```
      * 很多的动态规划都是自底向上的
      * 
