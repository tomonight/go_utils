package queue

type Queue interface {
	Add(d interface{}) bool
	Offer(d interface{}) bool
	Remove() interface{}
	Poll() interface{}
	Element() interface{}
	Peek() interface{}
}
