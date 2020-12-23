package list

import "unsafe"

type Object interface{}

type LikedList struct {
	data unsafe.Pointer
	len  int
}

type LikedNode struct {
	pre  *Node
	data interface{}
	next *Node
}

func (l *LikedNode) New(data Object) *LikedNode {

	return &LikedNode{data: data}
}

func (l *LikedNode) NewWithNext(data Object, ln *Node) *LikedNode {

	return &LikedNode{data: data, next: ln}
}

func (l *LikedNode) Add(data Object) {

}
