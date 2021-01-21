package main

type IStack interface {
	pop() string
	push(string)
	isEmpty() bool
}

//implement by linkednode

//implement by array


type IQueue interface {
	enqueue(string)
	dequeue() string
	isEmpty() bool
	size() int
}
//iterators