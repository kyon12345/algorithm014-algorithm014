## 深度优先搜索和广度优先搜索

#### dfs不是只针对树结构的，而是针对递归状态树的一种算法

* 比如括号生成问题的每一步都可以看做是一个选择树

## 贪心算法

* 每一个选择都取当前的最优解
* 贪心最重要的是分析最优结果是否是每一步最优结果的迭代
* 可以从前往后,也可以从后往前,也可以在某种条件下作为辅助算法

## 二分查找

* 牢记模板
```
# Python
left, right = 0, len(array) - 1 
while left <= right: 
	  mid = (left + right) / 2 
	  if array[mid] == target: 
		    # find the target!! 
		    break or return result 
	  elif array[mid] < target: 
		    left = mid + 1 
	  else: 
		    right = mid - 1
```
* 单调的,可通过索引访问的,有上下界的
* 使用二分查找,寻找一个半有序数组中间无序的地方(log(n))

```
func findBreak(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	l,r := 0,len(nums) - 1

	for {
		mid := (l + r) / 2

		if  r - l == 1 {
			return l
		}

		if nums[l] < nums[r] {
			r = mid
		} else {
			l = mid
		}
	}
}
```


## homework
---
[岛屿数量](day1/200.岛屿数量.go)
[单词接龙](day1/127.单词接龙.go)
[跳跃游戏](day2/55.跳跃游戏.go)
[跳跃游戏Ⅱ](day2/45.跳跃游戏-ii.go)