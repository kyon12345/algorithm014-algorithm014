## 字典树和并查集

---

### 字典树

实现

```
package main
// @lc code=start
type Trie struct {
	next [26]*Trie
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = &Trie{}
		}
		node = node.next[v]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		if node = node.next[v - 'a'];node == nil {
			return false
		}
	}

	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		if node = node.next[v - 'a'];node == nil {
			return false
		}
	}
	return true
}
```

-   字典树的每一个节点不是完整的单词
-   核心思想是利用空间换时间,利用字符串的公共前缀来降低查询的复杂度

### 并查集

实现

```
type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

// 路径压缩后 查询为O(1)
func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
```

-   并查集比较常见的实现方式是 parent[i] = i
-   通常用来解决合并连接的集合的问题
-   对于二维数组通常要做处理为一维数组
    -   n = x \* col + y

# 高级搜索

---

## 剪枝

## 双向 BFS

## 启发式搜索

-   每一次搜索都优先搜索 h(q)更大的地方,h 是估值函数
-   比如解数独问题,我们在搜索每一步的时候就先搜索数字较多的 Box,来达到优化的目的
-   这些搜索是可以叠加的,因为基本都是基于 BFS 来进行的,比如双向 BFS + A\*
-   x = n % col ,y = n /col manhatonDis = abs(xi - x) + abs(yi - y)

# 红黑树和 AVL 树

---

-   AVL 是更加严格平衡的二叉树
-   平衡二叉树产生的原因是为了避免树结构应为不平衡退化为链表
-   AVL 每个节点存储 balance factor = {1,0,-1}
-   通过旋转操作来保证树的平衡
    -   所以插入需要频繁的调整
-   红黑树的特点是从根节点到叶子的最长可能路径不能大于最短可能路径的两倍

## 对比

-   AVL 比红黑树的查询要快,因为他更加严格的平衡
-   红黑树的插入和删除操作更快,应为相对宽松的平衡
-   AVL 在每一个节点存储额外的信息,因此需要更多的存储空间
-   红黑树多用于各语言的库中比如 map,multimap,multisetin in c++,而 AVL 则多用于查询较多的数据库中
