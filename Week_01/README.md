学习笔记

## 作业

### 使用新的API改写queue的代码
```
deque.addFirst("a");
        deque.addFirst("b");
        deque.addFirst("c");
        System.out.println(deque);

        String str = deque.getFirst();
        System.out.println(str);
        System.out.println(deque);

        while (deque.size() > 0) {
            System.out.println(deque.removeFirst());
        }

        System.out.println(deque);
```
### Deque源码分析
这里以ArrayDequeue为例
[ArrayDeque源码分析](https://juejin.im/post/6844903833303253006)
这篇是java的,记录一下
golang Dequeue 有很多种实现方式
#### 示例代码
```
package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func main() {
	d := deque.New()
	for i := 0; i < 3; i++ {
		d.PushLeft(i)
	}
	// Pop out the deque contents and display them
	fmt.Println(d.PopLeft())
	fmt.Println(d.PopRight())
	fmt.Println(d.PopLeft())
}
```
#### 源码分析
```
package deque

// The size of a block of data
const blockSize = 4096

// Double ended queue data structure.
type Deque struct {
	leftIdx  int
	leftOff  int
	rightIdx int
	rightOff int

	blocks [][]interface{}
	left   []interface{}
	right  []interface{}
}

// Creates a new, empty deque.
//通过一个二维切片来实现
//[nil,nii .....]
//[nil,nil .....]
//....
//向right queue中添加元素时类似于向二维数组的横向添加元素,同时右移rightoff,为下一次添加做准备
//向left queue中添加则向二维数组的末端添加元素,左移leftoff
func New() *Deque {
	result := new(Deque)
	result.blocks = [][]interface{}{make([]interface{}, blockSize)}
	result.right = result.blocks[0]
	result.left = result.blocks[0]
	return result
}
```
### 学习总结
* 切忌和题目死磕到底,不要浪费时间,这样既折磨自己的耐心也对理解没有多大帮助,做不出来就果断看答案
* 没有思路的时候首先考虑暴力法求解,如果有暴力的方法就动手写,写不出来就参考别人的暴力法,这对理解很有帮助即使是得出的结果很慢很low也能理解题目的思路,求柱状图最大矩形的面积就是一个很好的例子,我们通过暴力法了解到这个问题的本质就是找左右矩形的最大边界,但是在枚举的过程中,我们可以发现寻找左边界的过程中有重复的计算,我们通过使用单向栈的方法来优化我们的暴力法,就能得到O(1)的复杂度
* 五遍刷题法,根据记忆曲线来刷题,很多解法看一遍是不够的,一看就会,一做就忘这种事情经常发生,所以还是得多练,只有形成了基本算法的肌肉记忆,才能挑战更高的层次