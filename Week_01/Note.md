##笔记
###时间复杂度和空间复杂度
时间复杂度是算法需要执行多少次,空间复杂度是一个算法在运行过程中占用的最大空间,主要取决于数组的长度和递归的深度

###递归问题
递归切忌傻递归,千万不要去人脑递归,只需要考虑当前层的逻辑
![求二叉树的最大深度](https://lylpic.oss-cn-beijing.aliyuncs.com/img/20200123232536.PNG "二叉树")

```
class Solution {
    public int maxDepth(TreeNode root) {
        //终止条件：当树为空时结束递归，并返回当前深度0
        if(root == null){
            return 0;
        }
        //root的左、右子树的最大深度
        int leftDepth = maxDepth(root.left);
        int rightDepth = maxDepth(root.right);
        //返回的是左右子树的最大深度+1
        return Math.max(leftDepth, rightDepth) + 1;
    }
}
```

同样的来思考两两交换链表中的节点这个问题

- **找终止条件**。 什么情况下递归终止？没得交换的时候，递归就终止了呗。因此当链表只剩一个节点或者没有节点的时候，自然递归就终止了。
- **找返回值。** 我们希望向上一级递归返回什么信息？由于我们的目的是两两交换链表中相邻的节点，因此自然希望交换给上一级递归的是已经完成交换处理，即已经处理好的链表。
- **本级递归应该做什么。** 结合第二步，看下图！由于只考虑本级递归，所以这个链表在我们眼里其实也就三个节点：head、head.next、已处理完的链表部分。而本级递归的任务也就是交换这3个节点中的前两个节点，就很easy了。

![两两交换链表](https://lylpic.oss-cn-beijing.aliyuncs.com/img/20200123232642.PNG)

```
class Solution {
    public ListNode swapPairs(ListNode head) {
      	//终止条件：链表只剩一个节点或者没节点了，没得交换了。返回的是已经处理好的链表
        if(head == null || head.next == null){
            return head;
        }
      	//一共三个节点:head, next, swapPairs(next.next)
      	//下面的任务便是交换这3个节点中的前两个节点
        ListNode next = head.next;
        head.next = swapPairs(next.next);
        next.next = head;
      	//根据第二步：返回给上一级的是当前已经完成交换后，即处理好了的链表部分
        return next;
    }
}
```


##go 源码分析 github.com/golang-collections/collections/stack
go 没有官方的stack结构,这里分析一下这个第三方的stack包,这个是通过链表来实现的栈结构
```
package stack

//通过链表来实现
type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}	
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
//interface{}用于返回未知类型的值
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
//Pop 和 Push 实际就是在做链表的增加和删除节点 
//查询复杂度 O(N) 修改复杂度O(1)
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}
```

