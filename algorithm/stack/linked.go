package stack

import "fmt"

type (
	LinkedStack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewLikedStack() *LinkedStack {
	return &LinkedStack{nil, 0}
}

// Return the number of items in the stack
func (this *LinkedStack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *LinkedStack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

func (this *LinkedStack) Show() {
	tmp := this.top
	fmt.Println("stack ==============")
	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.prev
	}
	fmt.Println("stack ==============")
}

// Pop the top item of the stack and return it
func (this *LinkedStack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *LinkedStack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}
