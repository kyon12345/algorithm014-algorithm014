学习笔记

##作业

###使用新的API改写queue的代码
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
###Deque源码分析
这里以ArrayDequeue为例
[ArrayDeque源码分析](https://juejin.im/post/6844903833303253006)
这篇是java的,记录一下有时间来复习
golang Dequeue 有很多种实现方式
####示例代码
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
####源码分析

