package stack

type stach interface {
	Pop() interface{}
	Push(interface{})
	Len() int
	Peek() interface{}
}
