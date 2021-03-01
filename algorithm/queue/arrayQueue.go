package queue

import "fmt"

type ArrayQueue struct {
	maxSize int
	front   int
	rear    int
	arr     []interface{}
}

var maxSize int

const defaultMaxSize = 1 << 12

func NewArrayQueue(size ...int) *ArrayQueue {
	maxSize = defaultMaxSize
	if len(size) != 0 {
		maxSize = size[0]
	}
	// use 1 memory to cache rear pointer
	return &ArrayQueue{maxSize: maxSize + 1, arr: []interface{}{}}
}

func (aq *ArrayQueue) IsEmpty() bool {
	return aq.rear == aq.front
}

func (aq *ArrayQueue) IsFull() bool {
	return (aq.rear+1)%aq.maxSize == aq.front
}

func (aq *ArrayQueue) Add(d interface{}) error {
	if aq.IsFull() {
		return fmt.Errorf("queue is full")
	}
	if len(aq.arr) < aq.maxSize {
		aq.arr = append(aq.arr, d)
	} else {
		aq.arr[aq.rear] = d
	}
	aq.rear = (aq.rear + 1) % aq.maxSize
	return nil
}

func (aq *ArrayQueue) Get() interface{} {
	if aq.IsEmpty() {
		return nil
	}
	data := aq.arr[aq.front]
	aq.front = (aq.front + 1) % aq.maxSize
	return data
}

func (aq *ArrayQueue) Size() int {
	return (aq.rear + aq.maxSize - aq.front) % aq.maxSize
}
