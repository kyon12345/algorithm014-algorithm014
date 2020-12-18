package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
	Age int
}

func (this People) Speak() {
	fmt.Println(this.Name)
}

func NewPeople() People {
	return People{"yon",10}
}

func main() {
	//p := &People{"yon",24}
	p := NewPeople()

	fmt.Println(reflect.TypeOf(&p))

	p.Speak()
}
