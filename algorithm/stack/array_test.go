package stack

import (
	"testing"

	"fmt"
)

func TestArrayStack_Peek(t *testing.T) {
	type fields struct {
		top   int
		stack []*arrayNode
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewArrayStack()
			s.Push(1)
			s.Push(2)
			s.Push(3)
			fmt.Println(s.Len())
			fmt.Println(s.Peek())
			fmt.Println(s.Pop())
			fmt.Println(s.Pop())
			fmt.Println(s.Pop())
			fmt.Println(s.Pop())
			fmt.Println(s.Pop())
		})
	}
}
