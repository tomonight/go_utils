package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue_Add(t *testing.T) {
	type fields struct {
		maxSize int
		front   int
		rear    int
		arr     []interface{}
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "success", args: args{d: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aq := NewArrayQueue(3)
			aq.Add(1)
			aq.Add(4)
			aq.Add(6)

			fmt.Println(aq.Get())
			fmt.Println(aq.Get())

			aq.Add(7)
			aq.Add(8)
			aq.Add(9)
			fmt.Println(aq.Get())
			fmt.Println(aq.Get())
			fmt.Println(aq.Get())
			fmt.Println(aq.Get())
		})
	}
}
