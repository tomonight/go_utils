package tree

import (
	"fmt"
	"testing"
)

func Test_pre(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root := &TreeNode{Val: 2}
	left := &TreeNode{Val: 3}
	right := &TreeNode{Val: 1}
	leftL := &TreeNode{Val: 3}
	leftR := &TreeNode{Val: 6}
	rightR := &TreeNode{Val: 7}
	left.Left = leftL
	left.Right = leftR
	right.Right = rightR
	root.Right = left
	root.Left = right
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success", args: args{root: root}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mid(tt.args.root)
			fmt.Println(got)
		})
	}
}
