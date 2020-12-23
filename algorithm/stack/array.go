package stack

type (
	ArrayStack struct {
		top   int
		stack []*arrayNode
	}
	arrayNode struct {
		data interface{}
	}
)

func NewArrayStack() *ArrayStack {
	return &ArrayStack{top: -1, stack: []*arrayNode{}}
}

func (s *ArrayStack) Pop() interface{} {
	if s.top == -1 {
		return nil
	}

	data := s.stack[s.top].data
	s.stack = s.stack[0:s.top]
	s.top--
	return data
}

func (s *ArrayStack) Len() int {
	return s.top + 1
}

func (s *ArrayStack) Push(data interface{}) {
	s.top++
	node := &arrayNode{data: data}
	s.stack = append(s.stack, node)
}

func (s *ArrayStack) Peek() interface{} {
	if s.top == -1 {
		return nil
	}

	return s.stack[s.top].data
}
