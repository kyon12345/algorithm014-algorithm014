## day1
### 哈希表,映射和集合

### 哈希表
通过**哈希函数**把关键码的值映射到表的某一个位置(index),表就是数组

#### 哈希碰撞
通过哈希函数得到了相同的值
工程中会把相同位置的值作为一个链表,但这样查询的时间复杂度会受到影响(O(1) - > O(n))

### 高级语言中的运用
map : key-value

set : 不重复元素的集合
