package tree

import (
	"fmt"
	"testing"
)

func Test_getBST(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "success", args: args{data: []int{10, 11, 7, 6, 8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newAVLTree(tt.args.data)
			//got.addBST(10)
			//deleteBST(got,nil,9)
			fmt.Println(got)
		})
	}
}
