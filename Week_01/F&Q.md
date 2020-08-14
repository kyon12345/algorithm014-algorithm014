```
for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
```
*基础,这两个循环是一样的吗?
不是,前者是先执行后判断,后者是先判断后执行
