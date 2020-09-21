学习笔记

DP:

* 分治子问题
* 状态数组定义
* DP方程


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

## DP的本质 就是分治 + 最优子结构

* 对于二维矩阵的问题,我们要考虑边界条件,有没有超出矩阵的范围
